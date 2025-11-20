package geecache

/*
用于实例化 lru 并封装 get 和 add 方法，添加互斥锁 mu 支持并发操作
*/

import (
	"geecache/lru"
	"sync"
)

// cache 为 lru.Cache 添加并发特性
type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

// add 添加节点，lru 为 nil 则先初始化 这叫延迟初始化(Lazy Initialization)，提高性能，并减少程序内存要求
func (c *cache) add(key string, value ByteView) {
	// 上锁解锁
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

// get 获取节点值，也加锁
func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		// 找到节点对应的值，先断言
		return v.(ByteView), true
	}
	return
}
