# 函数 function

> 具体代码见：[functions函数](../../../05_functions)

## 基本语法

```go
func 函数名(参数列表) 返回类型 {
// 函数体
return 值
}
```

例子：

```go
func add(a int, b int) int {
return a + b
}
```

---

## 参数写法

* **单个参数**
  `func f(x int)`
* **多个相同类型参数合并**
  `func f(x, y int, z string)`
* **参数与实参**

    * *param*（形参）：定义时写的
    * *arg*（实参）：调用时传的值

---

## 返回值

* **无返回值**

  ```go
  func hello() {
      fmt.Println("Hi")
  }
  ```
* **单返回值**

  ```go
  func square(x int) int {
      return x * x
  }
  ```
* **多返回值**

  ```go
  func divide(a, b int) (int, error) {
      if b == 0 {
          return 0, fmt.Errorf("除数不能为0")
      }
      return a / b, nil
  }
  ```
* **命名返回值**

  ```go
  func rect(w, h int) (area int, perimeter int) {
      area = w * h
      perimeter = 2 * (w + h)
      return // 直接返回
  }
  ```

---

## 可变参数

* **定义**

  ```go
  func sum(nums ...int) int {
      total := 0
      for _, v := range nums {
          total += v
      }
      return total
  }
  ```
* **调用**

    * 直接传：`sum(1,2,3)`
    * 切片展开：`sum(arr...)`

---

## 匿名函数 & 闭包

* **匿名函数**

  ```go
  f := func(msg string) {
      fmt.Println(msg)
  }
  f("Hello")
  ```
* **立即调用 (IIFE)**

  ```go
  func() {
      fmt.Println("Run now")
  }()
  ```
* **闭包**（函数记住外部变量）

  ```go
  func counter() func() int {
      x := 0
      return func() int {
          x++
          return x
      }
  }
  c := counter()
  fmt.Println(c()) // 1
  fmt.Println(c()) // 2
  ```

---

## 指针传参

* **值传递**（拷贝，不影响原值）
* **指针传递**（修改原始数据，避免大对象复制）

  ```go
  func setZero(x *int) {
      *x = 0
  }
  a := 5
  setZero(&a) // a = 0
  ```

---

## defer

* 延迟执行，**函数结束前执行（LIFO 顺序）**

  ```go
  func main() {
      defer fmt.Println("最后执行")
      fmt.Println("先执行")
  }
  ```

---

## 递归

* **函数调用自己**

  ```go
  func factorial(n int) int {
      if n == 0 {
          return 1
      }
      return n * factorial(n-1)
  }
  ```

---

## 回调函数

* **函数作为参数**

  ```go
  func process(data string, cb func(string)) {
      cb(data)
  }

  process("Go", func(s string) {
      fmt.Println("回调收到:", s)
  })
  ```

---

## 总结

* **形参定义，实参传值**
* **返回值可单可多，命名更直观**
* **...T 可变参数，切片需展开**
* **匿名 + 立即调用 = 临时逻辑**
* **闭包保留上下文，指针修改原对象**
* **defer 延迟清理，递归要有终止**
* **回调解耦逻辑，函数是一等公民**

