package lru

import "container/list"

// Cache 实现 LRU(Least Recently Used) 缓存，对于并发访问并不安全
type Cache struct {
	maxBytes  int64                         // 允许使用的最大内存，如果超过这个数，就会触发淘汰机制
	nBytes    int64                         // 当前已经使用的最大内存
	ll        *list.List                    // 标准库实现的双向链表
	cache     map[string]*list.Element      // O(1) 时间找到某个 key 对应的双向链表节点，从而进行获取，移动和删除操作
	OnEvicted func(key string, value Value) // 可以传一个函数，当数据被淘汰时调用，不需要可以是 nil
}

// entry 链表节点内容
type entry struct {
	key   string
	value Value
}

// Value 接口，缓存值必须能告诉自己占多少字节
type Value interface {
	Len() int
}

// Len 获取 Cache 中添加了多少个节点
func (c *Cache) Len() int {
	return c.ll.Len()
}

// New 用于构建 Cache 结构体
func New(maxBytes int64, OnEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nBytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}

// Get 查找功能，查找某个 key 对应的值
// 第一步是从字典中找到对应的双向链表的节点，第二步将该节点移动到队尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	element, ok := c.cache[key]
	if ok {
		// 找到对应节点，移动到链表头相当于最近使用过
		c.ll.MoveToFront(element)
		// 取出实际的键值对 entry 并返回，list.Element.Value 是 any，需要类型断言
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest 缓存淘汰，移除最少访问的节点（根据位置判断）
func (c *Cache) RemoveOldest() {
	// 获取队首也就是最不常访问的节点
	removeEle := c.ll.Back()
	if removeEle != nil {
		// 链表中删除
		c.ll.Remove(removeEle)
		kv := removeEle.Value.(*entry)
		// 从 c.cache 字典中删除该节点映射关系
		delete(c.cache, kv.key)
		// 更新当前所用的内存 c.nBytes
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		// 回调函数不为 nil，进行调用
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add 新增/修改 节点
func (c *Cache) Add(key string, value Value) {
	// 先从 cache 映射进行查找，找到就是修改，未找到就是新增
	if ele, ok := c.cache[key]; ok {
		// 找到了，修改
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		// 修改所占内存大小
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		// 修改 value
		kv.value = value
	} else {
		// 未找到，插入新元素到开头（这里使用 &entry 所以在删除和查找时查询出来的 value 需要使用 .(*entry) 进行断言）
		ele := c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	// 更新完判断当前 Cache 大小是否进行不常用节点的删除
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}
