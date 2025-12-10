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
	"sync"
	"sync/atomic"
)

// MagicNumber 魔数签名，用来标记这是 GeeRPC 的请求
const MagicNumber = 0x3bef5c

// Option 协议协商结构
// 客户端传来的整体报文结构大概是这样的：| Option(JSON) | Header(用Codec编码) | Body(用Codec编码) | Header2 | Body2 | ... |
// 在一次连接中只在最开始发一个 Option(JSON 编码)后面可以发很多次 Header + Body 请求
type Option struct {
	MagicNumber int        // 客户端要发来的值，如果跟服务端不一样直接拒绝处理
	CodecType   codec.Type // 告诉服务端使用哪种 Codec 对 Header 和 Body 进行解码编码
}

// DefaultOption 给客户的 / 示例代码用的默认配置
// 魔数固定，默认使用 Gob 进行编码解码
var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

// Server 目前只是个空 struct 后续塞入服务注册表等字段
type Server struct {
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
	server.serveCodec(f(conn))

}

var invalidRequest = struct{}{}

// serveCodec 请求主循环 + 并发处理
func (server *Server) serveCodec(cc codec.Codec) {
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
		go server.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()

}

// request 读取请求相关
type request struct {
	h            *codec.Header
	argv, replyV reflect.Value // 参数和返回值
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

	// TODO: 暂时不管请求参数类型，只认为是字符串，所以这里构造一个字符串容器，然后使用 ReadBody 将内容存入 req.argv
	req.argv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err: ", err)
	}
	return req, nil
}

// handleRequest 处理请求
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	// TODO 这里暂时只拼接打印参数，不做操作
	defer wg.Done()

	log.Println(req.h, req.argv.Elem())

	req.replyV = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))

	server.sendResponse(cc, req.h, req.replyV.Interface(), sending)

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

// 通过反射实现 service
type methodType struct {
	method    reflect.Method // 反射得到的 reflect.Method 本体(可以调用里面的 Func 字段)
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint64 // 统计被调用了多少次
}

// NumCalls 外部调用返回调用次数
func (m *methodType) NumCalls() uint64 {
	// 用 atomic.LoadUint64 来保证后面多 goroutine 统计时的原子性
	return atomic.LoadUint64(&m.numCalls)
}

// newArgv 根据方法入参类型 ArgType，创建一个合适的、可写的参数值 argv，后面好用反射/codec 去填充它
func (m *methodType) newArgv() reflect.Value {
	var argv reflect.Value
	if m.ArgType.Kind() == reflect.Ptr {
		// 方法参数是指针类型，返回对应数据的指针类型 Value
		argv = reflect.New(m.ArgType.Elem())
	} else {
		// 方法参数是值类型，返回对应数据的值类型 Value
		argv = reflect.New(m.ReplyType).Elem()
	}
	return argv
}
