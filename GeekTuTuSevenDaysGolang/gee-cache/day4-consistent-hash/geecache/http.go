package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
HTTP Server 服务端提供被其他节点访问的能力
*/

const defaultBasePath = "/_geecache/"

type HTTPPool struct {
	self     string // 当前节点自身地址， "http://localhost:8000"
	basePath string // HTTP 接口前缀，默认为 /_geecache/ 区分缓存系统的 API 请求
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

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

	// 6. 找到资源，通过原始字节数组（application/octet-stream）写入响应体
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}
