package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashFunc func(data []byte) uint32

// Map 一致性哈希算法的主数据结构
type Map struct {
	hash     HashFunc
	replicas int            // 虚拟节点倍数
	keys     []int          // 已排序的虚拟节点 hash 值
	hashMap  map[int]string // 从虚拟节点 hash 值映射到真实节点名称
}

// New 构造器，可以指定虚拟节点倍数和 hash 函数（默认 crc32.ChecksumIEEE）
func New(replicas int, fn HashFunc) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 添加真实节点/机器
func (m *Map) Add(keys ...string) {
	// 遍历真实节点 keys
	for _, key := range keys {
		// 每个 kye 生成对应个数的虚拟节点
		for i := 0; i < m.replicas; i++ {
			// 生成虚拟节点的 key 并计算 hash
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			// 放入 keys 数组
			m.keys = append(m.keys, hash)
			// 绑定虚拟节点对应的真实节点
			m.hashMap[hash] = key
		}
	}
	// 环上的哈希值按顺序排序，为 Get 使用二分查找做基础
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	// 计算 key 的 hash 值
	keyHash := int(m.hash([]byte(key)))
	// 二分查找找到到第一个 >=keyHash 的虚拟节点下标，如果 keyHash 比所有节点都大返回 len(m.keys)
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= keyHash
	})
	// 取出虚拟地址对应环上的真实地址，使用取余，如果 idx=len(m.keys) 则匹配到第一个元素也就是环的起点
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
