package main

import "fmt"

type Counter int

func (c *Counter) Add(n int) {
	*c += Counter(n)
}

func main() {
	/*
		方法值就是方法绑定到特定实例之后得到的可调用函数
		可以理解为闭包
		f := func(n int) {
			c.Add(n)
		}
		多个实例可以单独绑定方法值不会冲突
	*/
	var c1 Counter
	c2 := Counter(20)
	f1 := c1.Add // 方法值：绑定到实例 c 上
	f2 := c2.Add
	f1(10) // 此时这里就相当于 c.Add(10)
	f2(10)
	fmt.Println(c1) // 10
	fmt.Println(c2) // 30

}
