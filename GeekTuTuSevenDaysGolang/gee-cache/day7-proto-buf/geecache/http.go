package geecache

import (
	"fmt"
	"geecache/consistenthash"
	pb "geecache/geecachepb/geecachepb"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
)

/*
HTTP Server 服务端提供被其他节点访问的能力
*/

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50 // 虚拟节点数量，给一致性哈希用的
)

type HTTPPool struct {
	self        string                 // 当前节点自身地址， "http://localhost:8000"
	basePath    string                 // HTTP 接口前缀，默认为 /_geecache/ 区分缓存系统的 API 请求
	mu          sync.Mutex             // 保护 peers 和 httpGetter
	peers       *consistenthash.Map    // 一致哈希性结构
	httpGetters map[string]*httpGetter // 节点地址字符串，一台远程节点 = 一个 httpGetter 客户端
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

// Set 注册所有节点
func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	// 每个节点准备一个 HTTP 客户端
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

// PickPeer 实现接口，按 key 选出对应的节点
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

// 译期接口实现检查
var _ PeerPicker = (*HTTPPool)(nil)

func (p *HTTPPool) Log(format string, v ...any) {
	log.Printf("[Serve %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1. 判断 url 是否匹配 basePath 前缀
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpect path: " + r.URL.Path)
	}
	p.Log("%s %s", r.Method, r.URL.Path)

	// 2. 去除前缀后拆分 group 和 key
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// 3. 切分 groupName 和 key
	groupName := parts[0]
	key := parts[1]

	// 4. 根据 groupName 进行匹配
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no match group"+groupName, http.StatusNotFound)
		return
	}

	// 5. 根据 key 取 view
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	body, err := proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 6. 找到资源，通过原始字节数组（application/octet-stream）写入响应体
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
}

/*
Http 客户端
*/

type httpGetter struct {
	baseURL string // 将要访问的远程节点的地址
}

// Get 获取返回值，并转换为 []bytes 类型
func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(in.GetGroup()),
		url.QueryEscape(in.GetKey()),
	)
	res, err := http.Get(u)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}

	return nil

}

// 编译期接口检查 *httpGetter 是否完整实现了 PeerGetter 接口
var _ PeerGetter = (*httpGetter)(nil)
