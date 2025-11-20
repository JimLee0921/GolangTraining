package singleflight

import "sync"

/*
解决同一时候太多请求分到一个节点可能导致缓存击穿的问题
*/

// call 结构体表示一次正在进行中的调用，比如一次真正的加载（比如从 DB 拉数据）对应一个 call
type call struct {
	wg  sync.WaitGroup // 本次调用完成的信号器
	val any            // 调用返回值
	err error          // 错误信息
}

// Group 总控，表示某个 key 当前是否又正在执行中的 call
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (any, error)) (any, error) {
	g.mu.Lock()
	// 延迟初始化
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	// 判断当前 key 的调用是否已经存在
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	// 创建一次性 call
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	// 真正去 DB / 远端拉取数据，只执行一次
	c.val, c.err = fn()
	c.wg.Done() // wg 计数-1，所有等待者继续执行
	// 请理资源（只针对于短时间内的多个请求访问同一个key）
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
