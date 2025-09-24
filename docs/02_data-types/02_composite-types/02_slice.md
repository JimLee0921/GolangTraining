# 切片 slice

> 具体代码见：[slice切片](../../../02_data-type/03_composite-types/02_slice)

## 核心概念

- 切片（slice）是基于数组的轻量级视图，内部只保存指向底层数组的指针、长度（len）、容量（cap）
- 切片长度可变，元素可以重复，常用于动态序列处理，是 Go 中比数组更常用的容器形式

## 语法要点

### 创建方式

```go
var empty []int // 声明空切片，默认为 nil
var literal = []string{"Go"}     // 字面量初始化
nums := []int{1, 2, 3}           // 短变量声明
sized := make([]int, 3, 5) // 使用 make 分配 len=3, cap=5 的切片
buffer := make([]byte, 0, 8) // 创建空切片并预留容量
```

### 切片表达式

```go
source := []int{0, 1, 2, 3, 4, 5}
sub := source[1:4] // [1 2 3]  len=3 cap=5
full := source[1:4:4] // len=3 cap=3，限制容量避免影响原切片
```

- 普通表达式 `s[low:high]`：结果长度等于 `high-low`，容量等于 `cap(s)-low`
- 全切片表达式 `s[low:high:max]`：长度等于 `high-low`，容量等于 `max-low`，常用来控制 append 的影响范围

## 基础操作

```go
names := []string{"Jim", "Jane", "Tom"}
fmt.Println(names[0])        // 访问元素
names[1] = "Jerry"           // 修改元素
fmt.Println(len(names)) // 获取长度
fmt.Println(cap(names)) // 获取容量
```

- 下标从 0 开始，越界访问会触发运行时 panic
- `len` 表示当前可访问元素数量，`cap` 表示从起始位置到底层数组末尾的空间

## 追加与删除

```go
nums := []int{1, 2}
nums = append(nums, 3, 4) // 尾部追加
nums = append([]int{0}, nums...) // 头部追加
nums = append(nums[:2], nums[3:]...) // 删除索引 2 的元素
```

- `append` 会在容量不足时分配新数组并返回新切片，务必接收返回值
- 删除元素常用切片拼接或 `copy` 将尾部内容前移，再缩短长度

## 拷贝与遍历

```go
src := []int{1, 2, 3}
clone := make([]int, len(src))
copy(clone, src) // 深拷贝

clone2 := append([]int(nil), src...) // 另一种克隆写法

for i := 0; i < len(src); i++ {
fmt.Println(i, src[i])
}
for idx, val := range clone {
fmt.Println(idx, val)
}
```

- 直接赋值只会复制切片头部信息，多个切片共享底层数组
- 使用 `copy` 或 `append([]T{}, src...)` 才能获得互不影响的副本

## 扩容行为

- 当 `append` 使长度超过容量时，运行时会分配更大的底层数组并复制旧数据
- 容量增长并非固定倍数，小容量通常倍增，容量较大时增量会逐渐减小（约 1.25 倍）
- 扩容后的切片与旧切片不再共享底层数组

## 使用提醒

- 在使用切片表达式时子切片与原切片共享底层数组，修改其内容会影响原数据
- 多协程访问时注意同步，避免在未锁定的情况下同时 `append` 或修改
- 切片可用于实现栈、队列、窗口等结构，灵活组合即可满足多数动态容器需求
- `slices`
  包中有很多更方便的切片方法可以参考 [slices-package-example](../../../02_data-type/03_composite-types/02_slice/03_slices-package-example)