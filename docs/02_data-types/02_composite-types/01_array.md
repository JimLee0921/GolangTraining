# 数组 array

> 具体代码见：[array数组](../../../02_data-type/06_composite-types/01_array)

## 核心概念

- 数组（array） 是一组固定长度、相同类型的元素序列
- Go 中的数组是 值类型（赋值或传参会 复制整个数组）

## 语法要点

### 创建声明

```go
var arr1 [5]int // 定义长度为 5 的 int 数组，默认值都是 0
arr2 := [3]string{"Go", "Java", "Python"} // 初始化时指定元素
arr3 := [...]int{1, 2, 3, 4} // 用 ... 让编译器推断长度
arr4 := [5]int{0: 99, 3: 42} // 通过下标指定部分初始值
```

### 常用操作

定义数据

```go
// 定义两个一个数组
numberArray := [10]int{0: 55, 2: 22, 6: 2}
stringArray := [...]string{"go", "python", "java", "c#"}
```

1. 访问数组下标获取指定下标的值（下标从 0 开始）
    ```go
    fmt.Println(numberArray[2]) // 22
    ```

2. 使用 len 获取数组长度
    ```go
    fmt.Printf("stringArray的长度为: %d\n", len(stringArray))
    ```

3. 通过下标修改数组的某一个值
    ```go
    numberArray[5] = 123
    fmt.Println(numberArray)
    ```

4. 遍历数组：用 for 循环或 for ... range
    ```go
    for i := 0; i < len(stringArray); i++ {
        fmt.Printf("stringArray的第 %d 个值为: %v\n", i+1, stringArray[i])
    }
    // 可配合 break
    for i, v := range numberArray {
        fmt.Printf("numberArray[%d]: %v\n", i, v)
        if i > 6 {
            break
        }
    }
    ```

### 多维数组

数组的元素也可以是数组，从而形成多维数组
定义数组的那些操作和访问修改等操作都可以用于多维数组的定义
> 注意事项
> * Go 的数组长度是 类型的一部分，所以 [2][3]int 和 [2][2]int 是不同类型。
> * 创建多位数组使用 ... 推断写法（只能推断最外层数组）
> * 多维数组的多维数组的每一个维度长度必须相同（简言就是只能为矩阵，矩阵中不能有空元素）

## 常见陷阱

- 数组长度不同即为不同类型，无法直接赋值。
- 传递大数组可能造成性能压力，必要时改用切片指针。
