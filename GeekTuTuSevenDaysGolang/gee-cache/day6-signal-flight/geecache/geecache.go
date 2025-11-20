package geecache

import (
	"fmt"
	"geecache/singleflight"
	"log"
	"sync"
)

var (
	mu     sync.RWMutex              // 全局锁
	groups = make(map[string]*Group) // 每创建一个 Group 加入 groups
)

/*
GeeCache 最核心数据结构，负责于用户的交互并控制缓存值存储和获取
*/

// Getter 接口，传入 key 加载数据
// Group 本身不关心数据存哪里，只要事先这个接口，可以从 数据库 / 文件 / HTTP 自定义查找
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc 实现 Getter 接口
type GetterFunc func(key string) ([]byte, error)

// Get 回调函数，参数是 key，返回 []byte
// 接口型函数，方便使用者在调用时既能够传入函数作为参数，也能够传入实现了该接口的结构体作为参数
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// Group 核心定义
type Group struct {
	name      string // 每个组的命名空间
	getter    Getter // 缓存未命中时获取源数据的回调(callback)，由用户定义
	mainCache cache  // 并发缓存
	peers     PeerPicker
	loader    *singleflight.Group // 使用 singleflight 确保同一时间多个相同 key 只发起一次请求
}

// RegisterPeers 将实现了 PeerPicker 接口的 HTTPPool 注入到 Group 中
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

// NewGroup 函数用于创建一个 Group 实例并加入 groups 组，支持并发
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	// 规定必须传入 getter
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			cacheBytes: cacheBytes,
		},
		loader: &singleflight.Group{},
	}
	// 加入 groups
	groups[name] = g
	return g
}

// GetGroup 根据 group name 从 groups 中获取对应的 group，同样加锁
func GetGroup(name string) *Group {
	// 不涉及变量的写操作，只读锁即可
	mu.RLock()
	g := groups[name]
	defer mu.RUnlock()
	return g
}

// Get Group 核心方法
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	// 优先从 mainCache 进行擦哈找
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}
	// 没有找到，使用 load 方法调用 getLocally（分布式场景下会调用 getFromPeer 从其他节点获取）
	return g.Load(key)
}

// Load 使用 PickPeer() 方法选择节点
// 若非本机节点，则调用 getFromPeer() 从远程获取
// 若是本机节点或失败，则回退到 getLocally()
func (g *Group) Load(key string) (value ByteView, err error) {
	// 使用 g.loader.Do() 进行包装
	viewi, err := g.loader.Do(key, func() (any, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err = g.GetFromPeer(peer, key); err == nil {
					return value, nil
				} else {
					log.Println("[GeeCache] failed to get from peer", err)
				}
			}
		}
		return g.getLocally(key)
	})
	if err == nil {
		return viewi.(ByteView), nil
	}
	return
}

// getLocally 从本地数据源加载并放入缓存
func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.popularCache(key, value)
	return value, nil
}

// GetFromPeer 实现 PeerGetter 接口的 httpGetter 从访问远程节点，获取缓存值
func (g *Group) GetFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

// popularCache 将 key - value 写入 mainCache 并返回 ByteView 给调用方
func (g *Group) popularCache(key string, value ByteView) {
	// 真正写入 LRU
	g.mainCache.add(key, value)
}
