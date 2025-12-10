package geerpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type Call struct {
	Seq           uint64     // 本次调用序号
	ServiceMethod string     // 调用的方法名 如"Foo.Sum"
	Args          any        // 参数
	Reply         any        // 返回值
	Error         error      // 错误信息
	Done          chan *Call // 用于通知调用结束（使用管道支持异步）
}

// 本次 RPC 调用结束，通知等待的任务来取结果
func (call *Call) done() {
	call.Done <- call
}

// Client 真正的 GeeRPC 客户端
type Client struct {
	cc       codec.Codec
	opt      *Option
	sending  sync.Mutex   // 保护发送
	header   codec.Header // 复用的 Header
	mu       sync.Mutex
	seq      uint64           // 请求序号自增
	pending  map[uint64]*Call // 未完成的 Call
	closing  bool             // 用户主动关闭(即调用 Close 方法)
	shutdown bool             // 发生错误被动关闭
}

var _ io.Closer = (*Client)(nil)

// ErrShutDown 统一错误值
var ErrShutDown = errors.New("connection is shut down")

// Close 主动关闭客户端
func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closing {
		return ErrShutDown
	}
	client.closing = true
	return client.cc.Close()
}

// IsAvailable 外部判断是否还能使用
func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

// registerCall 管理 pending 调用，注册新发出的 call
// 给这个 RPC 调用发一个排队号，并且记在 pending 队列里
func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closing || client.shutdown {
		return 0, ErrShutDown
	}

	call.Seq = client.seq
	client.pending[call.Seq] = call
	client.seq++
	return call.Seq, nil
}

// removeCall 管理 pending 调用，拿掉已经成功/失败的 call
// 接收到响应时（在 receive 里）：根据 Header.Seq 找到对应的 call
// 发送请求失败时（send 里）：把挂在 pending 的 call 拿出来标错误
func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

// terminateCalls 管理 pending 调用，出现致命错误时，终止所有未完成调用
func (client *Client) terminateCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()

	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

// 创建 GeeRPC 客户端相关

func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		// 服务端先写的 Header 再写的 Body 客户端读取也需要先读取 Header，如果读取 Header 失败直接跳出循环
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		// 通过 Header 中的 seq 获取到对应的 call，先从 map 中移除并获取
		call := client.removeCall(h.Seq)
		switch {
		// 情况1: call 为空，说明没找到，但是服务器还是返回了这个响应，客户端需要把流都干净，也就是把Body读掉
		case call == nil:
			err = client.cc.ReadBody(nil)
		// 情况2: 服务端返回了 Error 字段，说明客户端已经无需读取 Body了，还是传入 nil 把流都干净，然后调用 call.done() 环境在这个 channel 上的用户协程
		case h.Error != "":
			call.Error = fmt.Errorf("rpc client error: %s", err)
			err = client.cc.ReadBody(nil)
			call.done()
		// 情况3: 正常返回，使用用户传入的 reply 指针读取 body
		default:
			err = client.cc.ReadBody(call.Reply)
			// 如果出错进行标记并调用 done
			if err != nil {
				call.Error = errors.New("reading body" + err.Error())
			}
			call.done()
		}
	}
	// 总清算，如果for循环推出说明存在 err ，关闭所有连接，处理所有未完成的 call
	client.terminateCalls(err)
}

// NewClient 外部使用方法，创建 Client
func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	// 通过传入的 opt 找到对应的构造函数
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error: ", err)
		return nil, err
	}
	// Option 协商阶段
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil
}

// newClientCodex 真正初始化 Client
func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1, // 从1开始，0为无效值
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	// 创建 Client 时直接启动 receive 在后台进行响应读取
	go client.receive()
	return client
}

// 配置 Option 使用可变参数，不传入使用默认配置，支持最多传入一个
func parseOptions(opts ...*Option) (*Option, error) {
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

// Dial 连接 RPC 服务，真正对外暴露的创建接口
func Dial(network, address string, opts ...*Option) (*Client, error) {
	return dialTimeout(NewClient, network, address, opts...)
}

// 实现发送请求相关

// send 串行的发送完整请求
func (client *Client) send(call *Call) {
	// 互斥锁保整包的完整性
	client.sending.Lock()
	defer client.sending.Unlock()

	// 先注册 call 放入 pending 如果注册出错就写入错误并关闭 call
	seq, err := client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	// 准备请求头和参数
	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	// 写入失败进行请理
	if err := client.cc.Write(&client.header, call.Args); err != nil {
		call := client.removeCall(seq)
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

// Go 异步调用接口，传入要调用的方法，参数和返回值指针
func (client *Client) Go(serviceMethod string, args, reply any, done chan *Call) *Call {
	// 如果不传入 done 会自动创建一个带缓冲的，如果传入 done 容量为 0 (无缓冲)直接 panic 因为会导致 call.done() 阻塞
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	// 真正发出请求
	client.send(call)
	return call
}

// Call 是对 Go 同步调用的包装，加上 context 支持超时处理
func (client *Client) Call(ctx context.Context, serviceMethod string, args, reply any) error {
	// 启动异步调用
	call := client.Go(serviceMethod, args, reply, make(chan *Call, 1))
	select {
	// 超时或取消直接移除调用并返回错误信息
	case <-ctx.Done():
		client.removeCall(call.Seq)
		return errors.New("rpc client: call failed:" + ctx.Err().Error())
	// RPC 调用先完成，receive 会把结果推送到 call.Done 这里进行接收并返回错误信息
	case call := <-call.Done:
		return call.Error
	}
}

type clientResult struct {
	client *Client
	err    error
}

type newClientFunc func(conn net.Conn, opt *Option) (client *Client, err error)

// dialTimeout 给 Dial 和 NewClient 套一层超时处理
func dialTimeout(f newClientFunc, network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}
	// 如果在指定时间内连不上，直接返回 error
	conn, err := net.DialTimeout(network, address, opt.ConnectTimeout)
	if err != nil {
		return nil, err
	}
	// 遇到错误关闭连接
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()
	// 在子 goroutine 中执行 NewClient
	ch := make(chan clientResult)
	go func() {
		client, err := f(conn, opt)
		ch <- clientResult{
			client: client,
			err:    err,
		}
	}()
	// 不限制 NewClient 时间，直接返回结果
	if opt.ConnectTimeout == 0 {
		result := <-ch
		return result.client, result.err
	}
	// 如果超时 client 返回 nil 并抛出错误
	select {
	case <-time.After(opt.ConnectTimeout):
		return nil, fmt.Errorf("rpc client: connect timeout: expect within %s", opt.ConnectTimeout)
	case result := <-ch:
		return result.client, result.err
	}
}
