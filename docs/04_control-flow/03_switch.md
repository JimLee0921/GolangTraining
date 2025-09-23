# `switch` 语句

> 具体代码见：[switch-statements](../../04_control-flow/03_switch-statements)

---

## 基本语法

### 表达式 `switch`（值匹配）

```go
x := 2
switch x {
case 1:
fmt.Println("x = 1")
case 2:
fmt.Println("x = 2")
default:
fmt.Println("other")
}
```

* Go **自动 break**，无需手写 `break`
* `default` 可选、位置不固定

### 条件 `switch`（相当于 if-else 链）

```go
score := 85
switch {
case score >= 90:
fmt.Println("优秀")
case score >= 60:
fmt.Println("及格")
default:
fmt.Println("不及格")
}
```

* 这类写法本质等同 `switch true { ... }`。
* 每个 `case` 必须是布尔表达式。

### 初始化语句

```go
switch y := len(name); {
case y == 0:
fmt.Println("empty")
case y < 5:
fmt.Println("short")
default:
fmt.Println("long")
}
```

* `y` 作用域仅在该 `switch` 代码块内。

### 单个 case 多个常量

```go
day := "Sat"
switch day {
case "Sat", "Sun":
fmt.Println("weekend")
default:
fmt.Println("weekday")
}
```

---

## `fallthrough`

如果 case 匹配到后遇到了 `fallthrough` 关键字就会直接穿透到下一个 case 语句（即使下一个 case 语句条件不匹配）

```go
n := 1
switch n {
case 1:
fmt.Println("one")
fallthrough
case 2:
fmt.Println("two") // 即使 n != 2 也会执行
}
```

注意：

* 只会落到**下一个** case，不能连跳多个。
* 用处较少；通常用“多个值一个 case”更清晰。

---

## 类型 `switch`接口断言

用于区分 `interface{}` 的动态类型：

```go
func printType(v interface{}) {
switch x := v.(type) {
case int:
fmt.Println("int:", x)
case string:
fmt.Println("string:", x)
case fmt.Stringer:
fmt.Println("Stringer:", x.String())
default:
fmt.Printf("unknown: %T\n", x)
}
}
```

要点：

* `switch x := v.(type) { ... }` 语法只能在 `type switch` 中使用
* `x` 在各个分支内是**已断言后的具体类型**

