package xclient

import (
	"context"
	. "geerpc"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

type XClient struct {
	d       Discovery
	mode    SelectMode
	opt     *Option
	mu      sync.Mutex
	clients map[string]*Client
}

// 编译器检查
var _ io.Closer = (*XClient)(nil)

func NewXClient(d Discovery, mode SelectMode, opt *Option) *XClient {
	return &XClient{
		d:       d,                        // 实例发现模块
		mode:    mode,                     // 负载均衡策略
		opt:     opt,                      // RPC 通讯选项
		clients: make(map[string]*Client), // 缓存的 Client 对象，按地址 key
	}
}

// Close 关闭方法，实现 io.Closer 接口
func (xc *XClient) Close() error {
	xc.mu.Lock()
	defer xc.mu.Unlock()
	// 关闭所有 Client
	for key, client := range xc.clients {
		_ = client.Close()
		delete(xc.clients, key)
	}
	return nil
}

func (xc *XClient) dial(rpcAddr string) (*Client, error) {
	xc.mu.Lock()
	defer xc.mu.Unlock()
	// 获取 client
	client, ok := xc.clients[rpcAddr]
	// 检查 client 是否存在并可用
	if ok && !client.IsAvailable() {
		_ = client.Close()
		delete(xc.clients, rpcAddr)
		client = nil
	}
	// client 不存在则重新 dial
	if client == nil {
		var err error
		client, err = XDial(rpcAddr, xc.opt)
		if err != nil {
			return nil, err
		}
		// 对新的 client 进行缓存
		xc.clients[rpcAddr] = client
	}
	// 返回可用的 client
	return client, nil
}

func (xc *XClient) call(rpcAddr string, ctx context.Context, serviceMethod string, args, reply any) error {
	client, err := xc.dial(rpcAddr)
	if err != nil {
		return err
	}
	return client.Call(ctx, serviceMethod, args, reply)
}

func (xc *XClient) Call(ctx context.Context, serviceMethod string, args, reply any) error {
	// 选择一个服务器实例地址
	rpcAddr, err := xc.d.Get(xc.mode)
	if err != nil {
		return err
	}
	// 调用具体的 RPC
	return xc.call(rpcAddr, ctx, serviceMethod, args, reply)
}

// BroadCast 对所有实例并发调用
func (xc *XClient) BroadCast(ctx context.Context, serviceMethod string, args, reply any) error {
	// 取出所有服务地址
	servers, err := xc.d.GetAll()
	if err != nil {
		return err
	}
	// 创建并发同步结构
	var mu sync.Mutex                      // 保护e 和 replyDone
	var wg sync.WaitGroup                  // 等待所有 goroutine 完成
	var e error                            // 存放第一个出现的错误
	replyDone := reply == nil              // 表示第一个成功回复是否已经写入 reply
	ctx, cancel := context.WithCancel(ctx) // 一旦某一个 rpc 请求成功或者失败，通知其它 goroutine 停止(节省资源)
	defer cancel()
	// 遍历所有实例，为每一个地址开启一个 goroutine
	for _, rpcAddr := range servers {
		wg.Add(1)
		go func(rpcAddr string) {
			defer wg.Done()
			// 对 reply 进行深拷贝
			// 因为每个实例的返回内容可能不同
			// 多 goroutine 不能共享同一个 reply 否则会竞争写入，所以每个实例都创建一个独立 reply 副本
			var clonedReply any
			if reply != nil {
				clonedReply = reflect.New(reflect.ValueOf(reply).Elem().Type()).Interface()
			}
			// 获取或创建对应实例的 client
			err := xc.call(rpcAddr, ctx, serviceMethod, args, clonedReply)
			mu.Lock()
			// 有实例出现错误
			if err != nil && e == nil {
				e = err
				cancel() // 若出现第一个错误，取消上下文，并通知其它 goroutine
			}
			// 某个实例成功且时第一个成功
			if err == nil && !replyDone {
				reflect.ValueOf(reply).Elem().Set(
					reflect.ValueOf(clonedReply).Elem())
				replyDone = true
			}
			mu.Unlock()
		}(rpcAddr)
	}
	wg.Wait()
	// 如果 e 是 nil 表示整次 broadcast 执行成功，如果非 nil 则返回第一次失败的错误
	return e
}

type GeeRegistryDiscovery struct {
	*MultiServersDiscovery               // 复用之前的负载均衡
	registry               string        // 注册中心地址
	timeout                time.Duration // 本地服务列表缓存有效期
	lastUpdate             time.Time     // 最近一次从 registry 更新时间
}

const defaultUpdateTimeout = time.Second * 10

func NewGeeRegistryDiscovery(registerAddr string, timeout time.Duration) *GeeRegistryDiscovery {
	if timeout == 0 {
		timeout = defaultUpdateTimeout
	}
	// timeout 设为 10 也就是最多是秒用一次旧的列表，10秒后主动去 registry 拉一次新列表
	d := &GeeRegistryDiscovery{
		MultiServersDiscovery: NewMultiServerDiscovery(make([]string, 0)),
		registry:              registerAddr,
		timeout:               timeout,
	}
	return d
}

// Update 本地更新+记录更新时间
func (d *GeeRegistryDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	d.lastUpdate = time.Now()
	return nil
}

// Refresh 过期后向注册中心 GET 一次
func (d *GeeRegistryDiscovery) Refresh() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.lastUpdate.Add(d.timeout).After(time.Now()) {
		return nil
	}
	log.Println("rpc registry: refresh servers from registry", d.registry)
	resp, err := http.Get(d.registry)
	if err != nil {
		log.Println("rpc registry refresh err:", err)
		return err
	}
	servers := strings.Split(resp.Header.Get("X-Geerpc-Servers"), ",")
	d.servers = make([]string, 0, len(servers))
	for _, server := range servers {
		if strings.TrimSpace(server) != "" {
			d.servers = append(d.servers, strings.TrimSpace(server))
		}
	}
	d.lastUpdate = time.Now()
	return nil
}

// Get / GetAll 确保调用前列表是最新的
func (d *GeeRegistryDiscovery) Get(mode SelectMode) (string, error) {
	if err := d.Refresh(); err != nil {
		return "", err
	}
	return d.MultiServersDiscovery.Get(mode)
}

func (d *GeeRegistryDiscovery) GetAll() ([]string, error) {
	if err := d.Refresh(); err != nil {
		return nil, err
	}
	return d.MultiServersDiscovery.GetAll()
}
