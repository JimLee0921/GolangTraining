# 布尔类型

> 具体代码见 [boolean](../../../02_data-type/05_basic-types/01_boolean/)

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

## 逻辑运算符

1. `!`（not）

    * 叫 **逻辑非（logical NOT）运算符**。
    * 一元运算符（作用于单个值），把 `true` 变成 `false`，`false` 变成 `true`。

2. `||`（or）

    * 叫 **逻辑或（logical OR）运算符**。
    * 二元运算符（作用于两个值），只要有一个为 `true`，结果就是 `true`。
    * 有 **短路求值**：左边已经 `true`，右边就不会再计算。

3. `&&`（and）

    * 叫 **逻辑与（logical AND）运算符**。
    * 二元运算符，只有两个值都为 `true`，结果才是 `true`。
    * 也有 **短路求值**：左边已经 `false`，右边就不会再计算。



