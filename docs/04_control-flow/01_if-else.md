# if-else 条件控制

> 具体代码见：[if-else](../../04_control-flow/01_if-else)

## 基础语法

```go
if 条件 {
// 当条件为 true 时执行
} else {
// 当条件为 false 时执行
}
```

* 条件必须是布尔表达式 (`true` / `false`)。
* 条件 **不需要括号** `()`。
* 语句块 **必须有花括号** `{}`，即使只有一行代码。

---

## if-else if-else 链调用

```go
if x < 0 {
fmt.Println("负数")
} else if x == 0 {
fmt.Println("零")
} else {
fmt.Println("正数")
}
```

* 多条件判断时按顺序匹配，命中一个后就不再继续判断。

---

## 初始化语句

Go 的 `if` 可以在条件前加一个**初始化语句**，并用分号隔开：

```go
if v := 10; v > 5 {
fmt.Println("v 大于 5")
} else {
fmt.Println("v 小于等于 5")
}
```

* 初始化语句里的变量作用域只在 `if-else` 语句块内有效。
* 可以一次初始化多个变量：
  ```go
  if a, b, c := 3, 5, "结果是a<b"; a < b {
      fmt.Println(c)
  }
  ```

---

## 注意事项

* **else 不能单独写初始化语句**，只能依赖 `if` 的初始化结果。
* `if` 的初始化语句只能写一条，但可以用 **多重赋值** 初始化多个变量。
* 变量作用域只在整个 `if-else` 链内部有效。


