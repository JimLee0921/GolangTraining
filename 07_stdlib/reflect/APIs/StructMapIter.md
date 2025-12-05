# reflect.MapIter

MapIter 是 Go 的 reflect 包提供的一个 迭代器类型，专门用来 安全、高效地遍历 map。
可以理解为反射版的 map 遍历器，它和 for range map 的内部逻辑一致。

```
type MapIter struct {}
```

## 存在意义

reflect 的 Value.MapKeys() 会：

- 返回 所有 key 的切片
- 会产生分配（alloc）
- map 很大时代价高
- map 在遍历期间修改会出问题

而 MapIter：

- 不会创建 keys 切片（更省内存）
- 允许一步步取 key/value
- 遍历顺序和 for range 完全一致
- 允许遍历过程中安全前进（但还是不建议修改 map）

它就是 reflect 级别的： `for k,v := range myMap {}` 的底层版本

## 核心方法

### 1. `Next()`

移动迭代器到 map 中的下一个键值，并返回是否还有元素。

```
func (iter *MapIter) Next() bool
```

**使用方式**

```
iter := v.MapRange()

for iter.Next() {        // 必须先 Next()
    k := iter.Key()      // 再读 Key
    val := iter.Value()  // 再读 Value
}
```

- 第一次调用 `Next()` 时，定位到第一个元素
- 后续每次调用 `Next()`，前进到下一个 KV
- 返回 false 表示遍历结束
- 顺序与正常 for range map 完全一致（伪随机）

> 必须先 Next 再 Key() / Value()，否则会导致 panic

### 2. `Key()`

在当前迭代位置返回 key，必须在 `Next()`返回 true 后调用，否则 panic。

```
func (iter *MapIter) Key() Value
```

- 返回的只是一个 只读 Value，不能 Set，不能修改
- 类型永远等于原 map 的 key 类型
- 底层是从 bucket entry 直接生成的 Value，不复制，不分配

### 3. `Value()`

返回当前元素的 Value，同样必须在 `Next()` 之后进行调用

```
func (iter *MapIter) Value() Value
```

- 同样是只读的 `reflect.Value` 如果要修改 value，需要使用：`v.SetMapIndex(k, newVal)`
- 因为 map 的 value 是不可寻址的（unaddressable），不能通过 `Value()` 得到的对象直接修改

### 4. `Reset()`

把一个 MapIter 重置到新的 map（或同一个 map 重新遍历）

```
func (iter *MapIter) Reset(v Value)
```

- 把内部的 bucket 指针、状态复位，显式绑定一个新的 map
- 内部会重新调用 runtime 的 `mapiterinit()`

**使用场景**

- 想复用 MapIter 避免重新分配（微优化）
- 一个函数里要多轮过滤某个 map
- 对同一个 map 做多轮操作，但不想每次都创建新的迭代器对象

## 和 MapKeys / MapIndex 差别

| 功能       | MapKeys        | MapIndex | MapRange(MapIter) |
|----------|----------------|----------|-------------------|
| 遍历 map   | ❌ 不能           | ❌ 不能     | ✔ 可以              |
| 获取所有 key | ✔ 返回切片         | ❌        | ❌                 |
| 避免内存分配   | ❌ 会分配切片        | ✔        | ✔                 |
| 遍历顺序     | 随 MapKeys 切片变化 | -        | ✔ 和 for range 一致  |
| 性能       | 慢（会分配）         | 中        | 快（最佳）             |
