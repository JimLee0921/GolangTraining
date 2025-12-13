package opeartion

import "sync"

// SafeMap 安全 Map
type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

// NewSafeMap 构造函数
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]int),
	}
}

// Get 方法
func (m *SafeMap) Get(key int) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.m[key]
	return v, ok
}

// Set 方法
func (m *SafeMap) Set(key, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}
