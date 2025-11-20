package geecache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	// f.Get() 也就是在调用上面 GetterFunc 定义的匿名回调函数
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Error("callback failed")
	}
}

func TestGroup_Get(t *testing.T) {
	// 使用 map 模拟真实数据库
	var db = map[string]string{
		"Tom":   "630",
		"Jim":   "555",
		"Bruce": "123",
	}

	// 定义 loadCounts map 统计每个 key 被数据库访问了多少次
	// 正常情况下应该第一次查找时
	loadCounts := make(map[string]int, len(db))
	// 创建 Group 并传入 GetterFunc（用户自己传入的查询逻辑）
	// 第一次访问，loadCounts 会递增
	// 第二次访问，如果缓存正确，则不会再调用 Getter ，loadCounts 不会变
	gee := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			// 从 db map 中取值
			if v, ok := db[key]; ok {
				// 计数器 loadCounts[key]++
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key] += 1
				// 返回 []byte(value)
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		},
	))
	// 遍历db 测试每一个 key
	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed to get value of %s", k)
		}
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}
}
