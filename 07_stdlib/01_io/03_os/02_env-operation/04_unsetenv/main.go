package main

import "os"

func main() {
	// os.Unsetenv 删除环境变量
	err := os.Unsetenv("MODE")
	if err != nil {
		panic(err)
	}
}
