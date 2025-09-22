# 布尔类型

## 核心概念

- 在 Go 中，布尔类型使用关键字 bool 表示
- 取值只有 true 或者 false，零值是 false，适合作为状态位或条件标记
- 最基础的逻辑值，常用于条件判断、循环控制

## 语法要点

- 不能与其它类型混用，因为 Go 是强类型语言，bool 不能隐式转为整数（C语言中 `true=1`，`false=0`）
    ```go
    var b bool = true
	// fmt.Println(int(b))  会爆出编译错误
    ```
- 如果声明但未赋值，布尔变量的默认值就是 false
    ```go
    var a bool // 声明布尔变量不赋值，默认值为 false
    fmt.Println(a)  // 输出 false
    
    b := true     // 简短声明自动判断类型
    fmt.Println(b)  // 输出 true
    ```
- 常用于条件判断控制结构：if、for 和 switch 等

## 实用示例

> 见 [boolean/main.go](../../../02_data-type/02_basic-types/boolean/main.go)

