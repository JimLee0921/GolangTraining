package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})
	// 添加真实节点，会生成虚拟节点 keys 为 [2, 4, 6, 12, 14, 16, 22, 24, 26]
	hash.Add("6", "4", "2")
	// 测试生成的虚拟节点对照的真实节点 cases
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// 添加新的真实节点 8，keys 更新为 [2, 4, 6, 8, 12, 14, 16, 18, 22, 24, 26, 28]
	hash.Add("8")

	// 27 现在应该属于 8 了
	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}
