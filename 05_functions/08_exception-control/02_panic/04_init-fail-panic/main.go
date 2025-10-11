package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	/*
		init 阶段 panic 是合理的（配置缺失属致命错误）
		init() 方法不需要手动调用
	*/

}

func init() {
	p := filepath.Join("temp_files", "config.yaml")
	data, err := os.ReadFile(p)
	if err != nil {
		panic(fmt.Sprintf("failed to read config %q: %v", p, err))
	}
	fmt.Println("config loaded:", len(data))
}
