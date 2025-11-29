package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(int8(0))
	// 反射赋值前验证是否会溢出
	fmt.Println(t.OverflowInt(127)) // false (可以容纳)
	fmt.Println(t.OverflowInt(130)) // true  (溢出)

	// 自动转换防止 panic
	err := func() error {
		if t.OverflowInt(10) {
		}
		return errors.New("overflow int")
	}()
	fmt.Println(err)

	fmt.Println(reflect.TypeOf(float64(29)).OverflowFloat(123131)) // false
}
