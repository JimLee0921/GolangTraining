# sync.Map

`sync.Map` 是 Go 标准库里针对特定并发访问模式优化的并发 map 实现

```
type Map struct {
	// contains filtered or unexported fields
}
```

Go 的原生 map 是天生不并发安全的：

```
m := make(map[string]int)

go func() {
    m["a"] = 1
}()

go func() {
    _ = m["a"]
}()
```

这段代码会导致资源竞争，虽然可以使用 `map + Mutex/RWMutex` 进行解决，但是会有较大的资源损耗

`sync.Map` 不是为了跟会更简单，而是为了在特定场景下更快，在读多、写少、key 生命周期长的场景中减少锁竞争，官方文档解释主要适用于以下两种情况：

1. key 在初始化后基本不再变化
2. 多个 goroutine 对 disjoint key 集合进行读写

> 读远多于写、很少 delete、key 稳定

## 核心方法

| 分组      | 方法                                     |
|---------|----------------------------------------|
| 基础 CRUD | `Load`, `Store`, `Delete`              |
| 原子复合操作  | `LoadOrStore`, `LoadAndDelete`, `Swap` |
| CAS 系列  | `CompareAndSwap`, `CompareAndDelete`   |
| 遍历      | `Range`                                |
| 清空      | `Clear`                                |

### 1. Load()

并发安全地读取 key 对应的值

```
func (m *Map) Load(key any) (value any, ok bool)
```

- `ok == true` 表示 key 存在，`ok == false` 表示 key 不存在
- 读操作是无锁快路径（读多写少时）
- 返回的 value 是 any 任意类型，需要类型断言

### 2. Store()

设置 key 的值（存在则覆盖，不存在则新增）

```
func (m *Map) Store(key, value any)
```

> 并发安全，写入操作，会走慢路径，不返回旧值

### 3. Delete()

删除 key，如果 key 不存在则什么都不做

```
func (m *Map) Delete(key any)
```

> 并发安全，不返回是否删除成功，删除频繁的场景不适合 `sync.Map`

### 4. LoadOrStore()

如果 key 已经存在，返回已有的值，否则存入 value 并返回它

```
func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)
```

- `loaded == true`：表示 key 已经存在，actual 是旧值
- `loaded == false`：新存入的值，`actual == value`

> 主要用于并发初始化、注册表、缓存填充

### 5. LoadAndDelete()

原子地读取并删除指定的 key

```
func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
```

> 主要用于只消费一次的数据，任务取走即删除

### 6. Swap

原子地设置新值，并返回旧值（如果存在）

```
func (m *Map) Swap(key, value any) (previous any, loaded bool)
```

- `existed == true`：old 是之前的值
- `existed == false`：之前不存在

> Store 不返回旧值，Swap 会返回旧值

### 7. CompareAndSwap

仅当当前值 == old 时，才会把它替换为 new

```
func (m *Map) CompareAndSwap(key, old, new any) (swapped bool)
```

- `ok == true`：替换成功
- `ok == false`：key 不存在，或值不等于 old

**注意事项**

- 比较使用 `==`，非深比较
- value 必须是可比较类型

### 8. CompareAndDelete()

仅当当前值 == old 时，才删除

```
func (m *Map) CompareAndDelete(key, old any) (deleted bool)
```

> 旧值必须时可比较类型，如果映射中不存在值的当前值，则返回 false

### 9. Range()

并发安全地遍历 map 中的元素

```
Range(f func(key, value any) bool)
```

- 弱一致性：可能看不到最新写入，可能看到即将被删除的项
- f 返回 false 会提前停止
- Range 只适合统计/观察/best-effort扫描

```
m.Range(func(k, v any) bool {
    fmt.Println(k, v)
    return true // 继续
})
```

### 10. Clear()

Go1.23+ 使用，删除 map 中所有键值对

```
func (m *Map) Clear()
```

- 并发安全，但代价不高
- 如果经常 Clear，`sync.Map` 可能不是好的选择