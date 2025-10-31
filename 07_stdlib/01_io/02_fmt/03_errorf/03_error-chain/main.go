package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	baseErr := errors.New("connect fail")
	wrapErr := fmt.Errorf("init fail: %w", baseErr)
	// 判断链中是否包含某个特定错误对象
	if errors.Is(wrapErr, baseErr) {
		fmt.Println("baseErr exists")
	}
	_, err := os.Open("not_exist.txt")
	if err != nil {
		fmt.Printf("error type: %T\n", err)
	}
	var pathErr *os.PathError
	// 判断链中是否存在某种错误类型，并提取出来
	if errors.As(err, &pathErr) {
		fmt.Println("error because file path")
		fmt.Println("operation：", pathErr.Op)
		fmt.Println("path：", pathErr.Path)
		fmt.Println("true error：", pathErr.Err)
	} else {
		fmt.Println("not file path error")
	}
}
