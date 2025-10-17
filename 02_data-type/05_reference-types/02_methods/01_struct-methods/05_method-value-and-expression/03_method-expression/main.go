package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	/*
		方法表达式是类型级别的方法引用，未绑定实例
	*/
	p1 := Point{1, 2}
	f := p1.Move // 方法值，绑定 p1
	f(3, 4)
	fmt.Println(p1)    // {4 6}
	g := (*Point).Move // 方法表达式，不绑定实例，绑定到类型
	g(&p1, 1, 2)       // 传递参数时需要手动传递实例
	fmt.Println(p1)
}

/*
| 特性     | 方法值（Method Value） | 方法表达式（Method Expression）   |
| ------ | ----------------- | -------------------------- |
| 是否绑定实例 | 已绑定             | 未绑定                      |
| 函数签名   | 不含接收者             | 含接收者参数                     |
| 调用方式   | `f(args...)`      | `g(receiver, args...)`     |
| 生成形式   | `v.Method`        | `T.Method` 或 `(*T).Method` |
| 常见用途   | 延迟调用、回调函数         | 泛型/函数式编程中作为函数引用            |
| 本质     | 闭包（封装了接收者）        | 普通函数                       |

*/
