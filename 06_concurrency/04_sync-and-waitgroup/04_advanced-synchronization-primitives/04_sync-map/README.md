## `sync.Map`

Go 语言原生 map 并不是线程安全的，对它进行并发读写操作的时候，需要加锁。`sync.map` 是一种并发安全的 map，在 Go 1.9 引入。
可以被多个 goroutine 同时读写，而不需要手动加锁。

适合：

* 高并发下的共享字典
* 键值读多写少的场景（例如缓存、注册表、在线用户表等）

---

## 基本用法

```
var m sync.Map

// 写入（或更新）
m.Store("name", "JimLee")

// 读取
if v, ok := m.Load("name"); ok {
    fmt.Println(v) // 输出: JimLee
}

// 删除
m.Delete("name")

// 遍历
m.Range(func (k, v any) bool {
    fmt.Println(k, v)
    return true // 返回 false 可中断遍历
})
```

---

## 常见方法

| 方法                             | 作用              | 说明                     |
|--------------------------------|-----------------|------------------------|
| `Store(key, value)`            | 存储或更新键值         | 若 key 已存在则覆盖           |
| `Load(key)`                    | 读取值             | 返回 (value, true/false) |
| `LoadOrStore(key, value)`      | 若存在则返回旧值，否则写入新值 | 常用于只初始化一次的场景           |
| `Delete(key)`                  | 删除键             | 并发安全                   |
| `Range(f func(k, v any) bool)` | 遍历              | 回调返回 false 可中断遍历       |

---

## 性能特征

| 特性          | 说明                     |
|-------------|------------------------|
| 并发安全        | 可同时读写，不需手动加锁           |
| 优化读性能       | 内部读写分离：读用无锁快路径、写走慢路径   |
| 不适合写多       | 高频写入性能低于 `map+RWMutex` |
| Range 无顺序保证 | 不能依赖遍历顺序               |
| 不可取地址或修改    | `m[key]++` 这种写法不可用     |


