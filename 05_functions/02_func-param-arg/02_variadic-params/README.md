## 可变参数

在 Go 中，可以定义一个函数，使它接受**任意数量的参数**。
语法是在类型前加上三个点 `...`：

```go
func sum(nums ...int) int {
total := 0
for _, n := range nums {
total += n
}
return total
}
```

调用时：

```go
fmt.Println(sum(1, 2, 3)) // 6
fmt.Println(sum(5, 10, 15)) // 30
```

---

### 原理

可变参数底部由 slice 切片实现，实则是一种语法糖语法糖，在函数体内部可变参数就是一个普通的 slice 切片，
`func sum(nums ...int)` 等价于 `func sum(nums []int)`

只是调用方式不同：

* 普通调用时，编译器自动把多个实参打包成 slice
* 若已有一个 slice，要传进去，需在末尾加 `...` 展开：

```go
a := []int{1, 2, 3}
sum(a...) // 展开 slice
```

### 调用方式

| 场景           | 写法                                    | 说明              |
|--------------|---------------------------------------|-----------------|
| 直接传多个参数      | `sum(1, 2, 3)`                        | 自动打包成 slice     |
| 已有 slice 要展开 | `sum(nums...)`                        | 用 `...` 展开      |
| 空参数          | `sum()`                               | 传空 slice        |
| 同时传其他参数      | `fmt.Printf(format string, a ...any)` | 前面是固定参数，后面是可变参数 |

### 多个参数 + 可变参数

在一个函数参数中有有多个参数时只能有一个是可变参数且可变参数必须放在最后一个参数位置上。

### 可变参数的类型特征

* 是一个切片，所以在函数内部可以 `append`、`len()`、`cap()`
* 可变参数可以是任意类型（包括 interface{} / any）
* 因为可变参数就是 slice，所以可以使用 append
    ```go
    func merge(base []int, adds ...int) []int {
    return append(base, adds...) // 展开 adds slice
    }
    fmt.Println(merge([]int{1, 2}, 3, 4, 5)) // [1 2 3 4 5]
    ```

### 常见标准库使用场景

| 函数                                           | 用法                          |
|----------------------------------------------|-----------------------------|
| `fmt.Println(a ...any)`                      | 接受任意类型参数并打印                 |
| `append(slice []T, elems ...T)`              | 在已有切片后追加任意数量元素              |
| `errors.Join(errs ...error)`                 | 将多个 error 合并成一个             |
| `strings.Join([]string, sep string)`（不是可变参数） | 注意区分：这个不是可变参数，而是显式 slice 参数 |

---

## 函数传递

如果要把一个可变参数函数当作参数传给另一个函数，也可以直接传 `...`：

```go
func wrapper(f func (...int), nums ...int) {
f(nums...) // 转发展开
}

func printer(nums ...int) {
fmt.Println(nums)
}

func main() {
wrapper(printer, 1, 2, 3) // [1 2 3]
}
```

### 总结

- 使用 `func f(args ...T)` 进行定义
- 传递方式： 多个参数或用 `slice...` 展开
- 在函数内部 表现为 `[]T`
- 只能出现在参数列表最后且只能有一个
- 常见用途：打印、日志、动态拼接、泛型包装      

