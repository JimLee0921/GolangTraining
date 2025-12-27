# maps 顶层函数

maps 包存放的都是 map 相关的通用操作函数，主要分为下面几类

## Map 转为 Seq 可迭代学列

这几个函数把 map 变为 `iter.Seq`/`iter.Seq2` （惰性迭代），便于与 iter 生态组合使用

- All：产出键值对序列，等价于遍历整个 map 的迭代器视图
- Keys：产出 key 序列
- Values：产出 value 序列

### `maps.Keys`

```
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K]
```

**参数**

- `[Map ~map]`：表示接收 map 或底层实 map 的自定义类型
- `K comparable`：map key 的语言层面约束（key需要是可比较类型， Go map 硬性要求）
- `V any`：对于 map value 值是没有限制的，可以是任意值

**返回值**

- `iter.Seq[K]`：使用方式等参考 iter 包函数迭代器

### `maps.Values`

```
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V]
```

**参数**

- `[Map ~map]`：表示接收 map 或底层实 map 的自定义类型
- `K comparable`：map key 的语言层面约束（key需要是可比较类型， Go map 硬性要求）
- `V any`：对于 map value 值是没有限制的，可以是任意值

**返回值**

- `iter.Seq[K]`：使用方式等参考 iter 包函数迭代器

### `maps.All`

返回 `(key, value)` 序列，每次 yield 一对

```
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V]
```

**参数**

- `[Map ~map]`：表示接收 map 或底层实 map 的自定义类型
- `K comparable`：map key 的语言层面约束（key需要是可比较类型， Go map 硬性要求）
- `V any`：对于 map value 值是没有限制的，可以是任意值

**返回值**

- `iter.Seq2[K, V]`：使用方式等参考 iter 包函数迭代器

## Seq 序列转为 Map

这些函数负责把 iter 产出的序列收集成 map，属于构造/物化

- Collect：把 `(K, V)` 序列汇聚成 `map[K]V`
- Insert：把 `(K,V)` 序列插入到 map

### `maps.Collect`

```
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
```

**参数**

- `seq iter.Seq2[K, V]`：一个键值对序列
    - 通常由 `maps.All` 生成或自己构造的 `iter.Seq2`
    - `K, V` 的泛型依旧是 K 需要可比较，V 是任意值（符合 Go map）

**返回值**

- `map[K]V`：一个新创建的 map，非 nil，包含 seq 中产生的所有 `(K, V)`

**注意事项**

1. 如果 seq 中同一个 key 出现多次，后一次会不断覆盖前一次，与普通 map 复制语义安全一致
2. 顺序不影响结果，但是影响覆盖行为，因为 map 本身是无序的，但是 seq 迭代顺序会影响哪个值最后写入
3. 和所有 map 操作一样，Collect 不会深拷贝

等价于：

```
res := make(map[K]V)
for k, v := range seq {
    res[k] = v
}
```

### `maps.Insert`

```
func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V])
```

**参数**

- `m Map`：目标 map，`[Map ~map]` 表示必须为 map 或底层实 map 的自定义类型
    - 必须非 nil
    - 会被原地修改
- `seq iter.Seq2[K, V]`：一个键值对序列
    - 通常由 `maps.All` 生成或自己构造的 `iter.Seq2`
    - `K, V` 的泛型依旧是 K 需要可比较，V 是任意值（符合 Go map）

**返回值**

无返回值，这是一个副作用函数，所有结果直接写入 m

**注意事项**

1. m 为 nil 会导致 panic
2. 覆盖语义与 map 复制一致
3. 与 Copy 对比，Insert 是从序列写入 map，而 Copy 是从 map 写入 map

等价于：

```
for k, v := range seq {
    m[K] = v
}
```

## Map 复制/合并

下面几个函数用于把一份 map 内容搬到另一份 map，常用于合并配置，聚合结果转移等

- Clone：返回一个新的 map（浅拷贝）
- Copy：把 src 的键值对复制到 dst（覆盖同名 key 的值，dst 必须是非 nil map，否则导致 panic）

### `maps.Clone`

返回一个原 Map 的浅拷贝

```
func Clone[M ~map[K]V, K comparable, V any](m M) M
```

**参数**

- `m M`：输入 map（或底层为 map 的自定义类型）

**返回值**

返回一个新的 map（类型仍然为 M），包含与 m 相同的键值对

**注意事项**

1. 浅拷贝：只复制键和值本身；若 V 是引用类型（slice/map/pointer/含引用字段的 struct），引用仍指向同一底层对象
2. 若 m == nil，通常返回 nil（而不是空 map），工程上要区分nil map和空 map语义。
3. 覆盖/冲突：不存在冲突问题，因为是新 map。

### `maps.Copy`

将一个 Map 中的所有键值对复制（合并）到另一个目标 Map 中

```
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)
```

**参数**

- `dst M1`：目标 map（被写入）
- `src M2`：源 map（被读取）
- 两者 key/value 类型必须兼容同一组 `K, V`（允许自定义 map 类型，只要底层是 `map[K]V`）

**返回值**

无返回值，原地修改 dst

**注意事项**

1. dst 必须非 nil
2. 如果 dst 已存在某 key，src 中同名 key 会进行覆盖
3. 同样是浅拷贝

## Map 删除/过滤

- DeleteFunc：按照指定个条件进行删除，属于原地修改 map 数组

### `maps.DeleteFunc`

根据定义的条件，批量删除 Map 中符合条件的键值对

```
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool)
```

**参数**

- `m M`：目标 map（会被删除元素）
- `del func(K, V) bool`：删除判断函数
    - 返回 true 删除该项
    - 返回 false 保留该项

**返回值**

无返回值，原地修改满足条件的键值对

**注意事项**

- m 为 nil 安全：对 nil map 执行 delete 不会 panic
- del 会被调用多次，会有性能损耗
- 删除顺序不保证（map 是无序的）

## Map 比较判断

下面两个主要是只读操作，进行 map 之间的比较

- Equal：用于 K,V 都可比较 (comparable) 时的直接判等
- EqualFunc：value 不可比较或需要自定义等价关系时使用

### `maps.Equal`

比较两个 Map 是否相等

```
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool
```

**参数**

- `m1, m2`：两个 map（或底层为 map 的自定义类型）
- 约束要求 K 和 V 都必须是 comparable 可比较类型（可以用 `==` 进行比较）

**返回值**

- `bool`：两个 map 是否键集合相同且每个 key 对于的 value 也相同

### `maps.EqualFunc`

允许自定义值（Value）在什么情况下算相等，主要用于当 Value 的类型不可直接比较（例如是 slice 切片）或想忽略某些差异时

```
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool
```

**参数**

- `m1, m2`：两个 map（或底层为 map 的自定义类型，Value 类型可以不同）
- `eq func(v1, v2) bool`：自定义 value 相等判定

**返回值**

- `bool`：两个 map 是否键集合相同，且对每个 key：`eq(m1[key], m2[key]) == true`

**注意事项**

1. 主要为了解决值不可比较问题
2. 键缺失仍然会判定不等
3. eq 必须是稳定等价关系