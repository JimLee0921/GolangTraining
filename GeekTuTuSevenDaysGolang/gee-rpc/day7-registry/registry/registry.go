package registry

import (
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

// 注册中心，客户端和服务端通过中策中心进行连接

type GeeRegistry struct {
	timeout time.Duration
	mu      sync.Mutex
	servers map[string]*ServerItem
}

type ServerItem struct {
	Addr  string    // 服务实例地址字符串
	start time.Time // 用于判断实例是否超时
}

const (
	defaultPath    = "/_geerpc_/registry"
	defaultTimeout = time.Minute * 5
)

func New(timeout time.Duration) *GeeRegistry {
	return &GeeRegistry{
		timeout: timeout,
		servers: make(map[string]*ServerItem),
	}
}

var DefaultGeeRegister = New(defaultTimeout)

// 注册服务/给服务续命
func (r *GeeRegistry) putServer(addr string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.servers[addr]
	// 没有找到服务就新注册
	if s == nil {
		r.servers[addr] = &ServerItem{
			Addr:  addr,
			start: time.Now(),
		}
		// 注册过了，就把时间戳刷新
	} else {
		s.start = time.Now()
	}
}

// 请理超时的服务并返回存活的服务列表
func (r *GeeRegistry) aliveServers() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var alive []string
	for addr, s := range r.servers {
		if r.timeout == 0 || s.start.Add(r.timeout).After(time.Now()) {
			alive = append(alive, addr)
		} else {
			delete(r.servers, addr)
		}
	}
	sort.Strings(alive)
	return alive
}

// 实现 HTTP 接口并只允许 GET 和 POST 请求
func (r *GeeRegistry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	// GET 请求响应头放一条自定义 header
	case "GET":
		w.Header().Set("X-Geerpc-Servers", strings.Join(r.aliveServers(), ","))
	// POST 请求就是服务端注册，需要把地址写在与注册中心允许的自定义请求头中
	case "POST":
		addr := req.Header.Get("X-Geerpc-Servers")
		if addr == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.putServer(addr)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleHTTP 暴露到默认的 HTTP mux
func (r *GeeRegistry) HandleHTTP(registryPath string) {
	http.Handle(registryPath, r)
	log.Println("rpc registry path: ", registryPath)
}

// HandleHTTP 方便外部直接调用默认注册中心
func HandleHTTP() {
	DefaultGeeRegister.HandleHTTP(defaultPath)
}

// HeartBeat 服务端定时向注册中心请求心跳
func HeartBeat(registry, addr string, duration time.Duration) {
	// 默认 timeout 为 5 分钟，这里默认检测心跳周期为 5-1 四分钟
	if duration == 0 {
		duration = defaultTimeout - time.Duration(1)*time.Minute
	}
	var err error
	// 服务器注册后立即发送一次心跳
	err = sendHeartBeat(registry, addr)
	// 再起一个子 goroutine 定时发送心跳
	go func() {
		t := time.NewTicker(duration)
		for err == nil {
			<-t.C
			err = sendHeartBeat(registry, addr)
		}
	}()
}

// 内部就是一个 HTTP POST 对应注册中心的 case "POST"
func sendHeartBeat(registry, addr string) error {
	log.Println(addr, "send heart beat to registry", registry)
	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", registry, nil)
	req.Header.Set("X-Geerpc-Server", addr)
	if _, err := httpClient.Do(req); err != nil {
		log.Println("rpc server: heart beat err: ", err)
		return err
	}
	return nil
}
