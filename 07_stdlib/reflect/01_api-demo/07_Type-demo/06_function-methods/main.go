package main

import (
	"fmt"
	"reflect"
)

func Foo(a int, b string, c ...float64) (int, error) {
	return 0, nil
}

func main() {
	t := reflect.TypeOf(Foo)
	fmt.Println(t.IsVariadic()) // true 是否为可变参数函数

	// 入参
	fmt.Println("参数个数:", t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		p := t.In(i)
		fmt.Printf("  In[%d]: %s (Kind=%s)\n", i, p, p.Kind())
	}

	// 出参
	fmt.Println("返回值个数:", t.NumOut())
	for i := 0; i < t.NumOut(); i++ {
		r := t.Out(i)
		fmt.Printf("  Out[%d]: %s (Kind=%s)\n", i, r, r.Kind())
	}
}
