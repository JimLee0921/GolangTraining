# 字符串

> 具体代码见 [string/main.go](../../../02_data-type/02_basic-types/string/main.go)

## 基本介绍

* 在 Go 中，字符串用 **`string`** 表示
* **字符串是不可变的**：一旦创建，内部的字节序列不能被修改
    * 安全性：避免并发程序里字符串被意外修改
    * 效率率：多个变量可以安全地共享同一份字符串数据
    * 简洁性：字符串可以当作值来拷贝、比较，而不必担心底层数据被改变
* 底层本质上是一个 **字节序列（UTF-8 编码）**

---

## 核心概念

- **声明与赋值**：
    ```go
    var s1 string = "Hello"
    s2 := "Go语言"
    fmt.Println(s1, s2) // Hello Go语言
    ```
- **零值为空字符串**：在定义字符串时如果没有指定赋值则为空字符串
    ```go
    var s string
    fmt.Println(s == "") // true
    ```
- **获取长度**：字符串可以通过 len 方法获取其长度
    ```go
    fmt.Println(len("Hello")) // 5
    fmt.Println(len("Go语言")) // 8（因为 UTF-8 编码下中文占 3 个字节）
    ```
- **索引和遍历**：字符串可以通过下标（从 0 开始）访问字节
  > 注意：索引取到的是 **字节**，不是字符（Unicode）遍历字符要用 `for range`
    ```go
    s := "Go语言"
    fmt.println(s[0]) // 71（G的字节值）
    fmt.Printf("%c\n", s[0]) // G（%c为字符占位符）
    // 遍历
    for i, r := range "Go语言" {
        fmt.Printf("%d: %c\n", i, r) // 也可以直接用 fmt.Println(i, string(c))
    }
    /*
    输出:
    0: G
    1: o
    2: 语
    5: 言
    */
    ```
- **字符串拼接**：
    ```go
    s1 := "Hello"
    s2 := "World"
    s3 := s1 + " " + s2
    fmt.Println(s3) // Hello World
    ```
- **多行字符串**：可以使用 **\` \`** 反引号定义多行字符串，内容会原样保存
    ```go
    text := `第一行
    第二行
        保留空格`
    fmt.Println(text)
    ```
- **类型转换**：字符串和数字之间的转换
    ```go
    import "strconv"
    
    i := 123
    s := strconv.Itoa(i) // int -> string "123"
    n, _ := strconv.Atoi(s) // string -> int 123
    ```
- **特殊类型**：`rune` 与字符。`byte`：表示单个字节，常用于英文字符。`rune`：表示单个 Unicode 字符（int32）。
    ```go
    ch1 := 'A'  // rune
    ch2 := '中' // rune
    fmt.Printf("%c %d\n", ch1, ch1) // A 65
    fmt.Printf("%c %d\n", ch2, ch2) // 中 20013
    ```

---

## 注意事项

* `string` 是不可变的字节序列。
* UTF-8 编码下，一个中文占 3 个字节。
* 遍历字符串时最好用 `for range`，保证不会拆坏 Unicode 字符（如果使用 `for index` 在存在中文的情况下遍历的还是字节而不是字符）
* 常见操作都在 `strings` 包和 `strconv` 包里。 见 [string/main.go](../../../02_data-type/02_basic-types/string/main.go)
* `string(i)` 会把整数 i 当成 Unicode 码点 转成对应的字符（rune -> UTF-8 字节序列 -> string）
* []rune(str) 可以把字符串解析成一组 Unicode 码点
* []byte(str) 则把字符串拆成 UTF-8 字节
