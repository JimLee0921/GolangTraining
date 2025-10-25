package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 获取所有环境变量
	envs := os.Environ()
	for _, env := range envs {
		fmt.Println(env)
	}
	// 使用 strings.SplitN 拆分键和值
	for _, env := range envs {
		pair := strings.SplitN(env, "=", 2)
		fmt.Println("Key:", pair[0], "Value:", pair[1])
	}
}
