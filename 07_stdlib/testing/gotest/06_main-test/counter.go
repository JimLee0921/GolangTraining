package counter

import "sync"

type Counter struct {
	mu sync.Mutex
	n  int
}

// NewCounter 初始化 Counter
func NewCounter(init int) *Counter {
	return &Counter{
		n: init,
	}
}

func (c *Counter) Inc() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
	return c.n
}

func (c *Counter) Add(delta int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n += delta
	return c.n
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}
