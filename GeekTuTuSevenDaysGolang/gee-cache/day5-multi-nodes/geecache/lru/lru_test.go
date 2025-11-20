package lru

import (
	"reflect"
	"testing"
)

// 自定义 String 类型实现 Len 方法
type String string

func (d String) Len() int {
	return len(d)
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("KeyOne", String("12345"))
	if v, ok := lru.Get("KeyOne"); !ok || string(v.(String)) != "12345" {
		t.Fatal("cache hit KeyOne=12345 failed")
	}
	if _, ok := lru.Get("KeyTwo"); ok {
		t.Fatal("cache miss KeyTwo failed")
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	// 测试内存超出设定值是否会触发删除无用节点
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	// 自定义容量只够 kv1 + kv2
	myCap := len(k1 + k2 + v1 + v2)
	lru := New(int64(myCap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	// 容量满了，默认应该把 kv1 删除
	lru.Add(k3, String(v3))

	if _, ok := lru.Get(k1); ok || lru.Len() != 2 {
		t.Fatalf("RemoveOldest %s failed", k1)
	}
}

func TestCache_OnEvicted(t *testing.T) {
	// 传入 OnEvicted 测试是否会触发
	keys := make([]string, 0)
	// 定义删除时的回调函数把删除的 key 追加到 keys
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	// kv1, kv2 大小都是一样的字节，都是10
	k1, k2 := "key1", "key2"
	v1, v2 := "value1", "value2"
	myCap := len(k1 + v1)
	lru := New(int64(myCap), callback)
	lru.Add(k1, String(v1))
	// 容量满了，默认应该把 kv1 删除并调用 callback 把 k1 存入 keys
	lru.Add(k2, String(v2))
	// 容量满了，默认应该把 kv2 删除并调用 callback 把 k1 存入 keys
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4")) // kv3 + kv4 不够 10，不会进行删除操作
	expect := []string{k1, k2}
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
