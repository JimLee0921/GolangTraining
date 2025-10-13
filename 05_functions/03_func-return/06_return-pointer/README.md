## 基本概念

函数不仅可以返回普通值，还可以返回某个值的 **指针**（即该值的内存地址），在 Go 语言中，**返回指针（returning a pointer）**
是一个非常常见且安全的做法，尤其在结构体或大型数据类型中。

```go
func newInt() *int {
x := 10
return &x
}
```

这里：

* `x` 是一个局部变量
* `&x` 取它的地址
* 函数返回 `*int` 类型（指向 `int` 的指针）

在 Go 中这是合法的，因为：

> Go 编译器会自动将该局部变量“逃逸”到堆上（heap），保证返回后仍然有效

## 结构体指针的返回

结构体返回指针是非常常见的模式：

```go
type Point struct {
X, Y int
}

func NewPoint(x, y int) *Point {
return &Point{X: x, Y: y}
}

func main() {
p := NewPoint(3, 4)
fmt.Println(p.X, p.Y) // 3 4
}
```

这种写法比返回整个结构体更高效，因为：

* 结构体可能很大
* 返回指针避免了值拷贝
* 还能共享修改同一个对象

## 与值返回的区别

| 返回方式     | 内存位置   | 是否拷贝 | 修改是否影响原值 |
|----------|--------|------|----------|
| **值返回**  | 拷贝新值   | 是    | 否        |
| **指针返回** | 引用同一内存 | 否    | 是        |

## 返回指针的典型用途

1. **构造函数模式**

   ```go
   func NewUser(name string) *User {
       return &User{Name: name}
   }
   ```
   - 方便在别的地方共享引用
    - 节省结构体拷贝
    - 特别常见于 ORM（比如 GORM）、序列化等场景

2. **可选返回（类似引用语义）**

   ```go
   func findEven(n int) *int {
       if n%2 == 0 {
           return &n
       }
       return nil
   }
   ```

3. **配合 JSON、数据库、ORM 等库**

    * `*int`, `*string`, `*bool` 常用于表示可空字段


