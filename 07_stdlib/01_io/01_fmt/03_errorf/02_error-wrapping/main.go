package main

import (
	"errors"
	"fmt"
)

func main() {
	baseErr := errors.New("connect fail")
	// 使用 %w 把一个错误包在另一个错误里面（也可以看作占位符，类型是error）
	// wrapErr 不仅是一个新的错误消息，还保留了原始错误
	wrapErr := fmt.Errorf("init fail: %w", baseErr)
	fmt.Println(baseErr)
	fmt.Println(wrapErr)
}
