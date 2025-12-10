package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"strings"
	"sync"
	"time"
)

// MagicNumber 魔数签名，用来标记这是 GeeRPC 的请求
const MagicNumber = 0x3bef5c

// Option 协议协商结构
// 客户端传来的整体报文结构大概是这样的：| Option(JSON) | Header(用Codec编码) | Body(用Codec编码) | Header2 | Body2 | ... |
// 在一次连接中只在最开始发一个 Option(JSON 编码)后面可以发很多次 Header + Body 请求
type Option struct {
	MagicNumber    int           // 客户端要发来的值，如果跟服务端不一样直接拒绝处理
	CodecType      codec.Type    // 告诉服务端使用哪种 Codec 对 Header 和 Body 进行解码编码
	ConnectTimeout time.Duration // 超时处理，0代表不做限制
	HandleTimeout  time.Duration
}

// DefaultOption 给客户的 / 示例代码用的默认配置
// 魔数固定，默认使用 Gob 进行编码解码
var DefaultOption = &Option{
	MagicNumber:    MagicNumber,
	CodecType:      codec.GobType,
	ConnectTimeout: time.Second + 10,
}

// Server 目前只是个空 struct 后续塞入服务注册表等字段
type Server struct {
	serviceMap sync.Map // 存储 service 映射
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

// Accept 接收新连接
// 一个连接对应一个 goroutine，后面在这个 goroutine 内部再处理此连接上的多个请求
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("rpc server: accept error: ", err)
			return
		}
		go server.ServeConn(conn)
	}
}

// Accept 包装一个顶层函数方便外部 geerpc.Accept(lis) 即可调用
func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

// ServeConn 先接收 Option 再选择 Codec
func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	// 确保连接关闭
	defer func() { _ = conn.Close() }()
	// 一个连接只读取一次 Option
	var opt Option

	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: option error", err)
		return
	}
	// 校验 MagicNumber
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	// 选择 Codec 构造函数
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	// 具体处理请求剩下的的 Headers+Bodies
	server.serveCodec(f(conn), &opt)

}

var invalidRequest = struct{}{}

// serveCodec 请求主循环 + 并发处理
func (server *Server) serveCodec(cc codec.Codec, opt *Option) {
	sending := new(sync.Mutex) // 确保每次响应是完整的
	wg := new(sync.WaitGroup)  // 等待所有请求处理完再关闭连接
	// 不断读取 request
	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break // 读取不到 Header 直接结束
			}
			// 存入 err 错误信息
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg, opt.HandleTimeout)
	}
	wg.Wait()
	_ = cc.Close()

}

// request 读取请求相关
type request struct {
	h            *codec.Header
	argv, replyV reflect.Value // 参数和返回值
	mtype        *methodType
	svc          *service
}

// readRequestHeader 读取请求头
func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && !errors.Is(err, io.ErrUnexpectedEOF) {
			log.Println("rpc server: read header error", err)
		}
		return nil, err
	}
	return &h, nil
}

// readRequest 读取请求
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{h: h}
	// 通过header查找service和method
	req.svc, req.mtype, err = server.findService(h.ServiceMethod)
	if err != nil {
		return req, err
	}
	// 生成参数值和返回值正确类型
	req.argv = req.mtype.newArgv()
	req.replyV = req.mtype.newReplyV()
	// body 解码时要保证是指针
	argvi := req.argv.Interface()
	if req.argv.Type().Kind() != reflect.Ptr {
		argvi = req.argv.Addr().Interface()
	}
	if err = cc.ReadBody(argvi); err != nil {
		log.Println("rpc server: read body err: ", err)
		return req, err
	}
	return req, nil
}

// handleRequest 处理请求，真正执行 RPC
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup, timeout time.Duration) {
	// 根据实际方法进行调用
	defer wg.Done()
	// 这里需要两个 channel 监控两件事，都设为空 struct 零开销
	// 1. call 是否已经执行完毕(无论成功还是失败)
	// 2. 响应是否已经写回客户端
	called := make(chan struct{}) // 已执行完方法（call 已返回）
	sent := make(chan struct{})   // 已写完响应（sendResponse 完成）
	// 启动业务执行 goroutine
	go func() {
		// 此处可能卡死，比如 SQL 查询慢，死循环，网络依赖停住
		err := req.svc.call(req.mtype, req.argv, req.replyV)
		// call 执行完毕
		called <- struct{}{}
		// 写回应
		if err != nil {
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			sent <- struct{}{}
			return
		}
		server.sendResponse(cc, req.h, req.replyV.Interface(), sending)
		// 回应发送完毕
		sent <- struct{}{}
	}()
	// 服务器不启用超时
	if timeout == 0 {
		<-called
		<-sent
		return
	}
	select {
	// 1. 超时情况处理
	case <-time.After(timeout):
		req.h.Error = fmt.Sprintf("rpc server: request handle timeout: expect within %s", timeout)
		server.sendResponse(cc, req.h, invalidRequest, sending)
	// 2.业务执行成功完成(收到 called 方法执行完毕，但是 sendResponse 还没保证写完)
	case <-called:
		<-sent
	}

}

func Register(rcvr any) error {
	return DefaultServer.Register(rcvr)
}

// Register 用户侧调用的API
func (server *Server) Register(rcvr any) error {
	// 自动把结构体实例转为 service
	s := newService(rcvr)
	// 按服务名存储对应的 service 如果已存在同名服务会报错
	if _, dup := server.serviceMap.LoadOrStore(s.name, s); dup {
		return errors.New("rpc: server already defined: " + s.name)
	}
	return nil
}

func (server *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {
	// 拆分服务名和方法名
	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		err = errors.New("rpc server: service/method request ill-formed: " + serviceMethod)
		return
	}
	serviceName, methodName := serviceMethod[:dot], serviceMethod[dot+1:]
	// 查找服务名是否已注册
	svci, ok := server.serviceMap.Load(serviceName)
	if !ok {
		err = errors.New("rpc server: can't find service: " + serviceName)
		return
	}
	// 查找对应方法
	svc = svci.(*service)
	mtype = svc.method[methodName]
	if mtype == nil {
		err = errors.New("rpc server: can't find method: " + methodName)
	}
	return
}

// sendResponse 发送响应相关
func (server *Server) sendResponse(
	cc codec.Codec,
	h *codec.Header,
	body any,
	sending *sync.Mutex,
) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error: ", err)
	}
}
