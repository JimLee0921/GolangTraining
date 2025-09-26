# 映射 map

> 具体代码见：[map映射](../../../02_data-type/06_composite-types/03_map)

## 基础概念

* `map[KeyType]ValueType`：键值对集合，底层基于哈希表实现。
* **无序**：遍历时键的顺序不固定。
* **零值**：未找到键时返回值类型的零值。

---

## 创建 map

```go
// 1. 字面量
m := map[string]int{"Alice": 23, "Bob": 30}

// 2. make 创建（常用）
m := make(map[string]int)

// 3. 指定容量（优化性能）
m := make(map[string]int, 100) 
```

---

## 基本操作

```go
m["Alice"] = 23 // 添加/更新
age := m["Alice"]  // 读取
delete(m, "Alice") // 删除

val, ok := m["Bob"] // 查询键是否存在
if ok {
fmt.Println("存在", val)
}
```

---

## 遍历

```go
for k, v := range m {
fmt.Println(k, v)
}
```

> 遍历顺序是随机的，不能依赖。

---

## 嵌套用法

* **map of slice**：`map[string][]int`
* **map of map**：`map[int]map[string]int`
* 使用前要先初始化内部结构，否则会 `panic`。

示例：

```go
buckets := make(map[int]map[string]int)
for i := 0; i < 12; i++ {
buckets[i] = make(map[string]int)
}
buckets[3]["hello"]++
```

---

## 与切片结合

* **桶计数**：`[]int` 作为桶，用 `buckets[i]++` 累加。
* **桶存元素**：`[][]string`，每个桶里保存单词列表。
* **桶存 map**：`map[int]map[string]int`，每个桶里存单词 → 计数。

---

## 哈希相关小例子

```go
// 最简单的哈希：取首字母
func hashBucket(word string, buckets int) int {
return int(word[0]) % buckets
}

// 改进版：累加所有字符
func hashBucket(word string, buckets int) int {
sum := 0
for _, r := range word {
sum += int(r)
}
return sum % buckets
}
```

---

## 常见注意点

1. **map 不是并发安全的**，并发写要用 `sync.Map` 或加锁。
2. **map 的值是引用类型**，传参时是引用拷贝。
3. 不能对 map 取地址：`&m["key"]` 是非法的。
4. 不能直接比较两个 map（只能与 `nil` 比较）。

---

## 典型应用

* 词频统计：`map[string]int`
* 去重集合：`map[string]struct{}`
* 分桶/哈希表：`map[int][]T` 或 `map[int]map[K]V`

