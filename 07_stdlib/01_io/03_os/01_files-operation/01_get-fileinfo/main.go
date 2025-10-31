package main

import (
	"fmt"
	"os"
)

func main() {
	path := "temp_files/log.txt"
	/*
		os.Stat：如果 path 是符号链接，会跟随到目标再返回信息
		os.Lstat：拿的就是链接本身的信息，在需要辨别这是否为 symlink 时用
	*/
	info, err := os.Stat(path)
	if err != nil {
		// 判断是否存在
		if os.IsNotExist(err) {
			// 不存在，可以创建/忽略/报错
			fmt.Println(path, "not existing")
			return
		}
		// 其他错误 （可能全选/路径无效等）
		fmt.Println("stat error", err)
		return
	}
	/*
		fileInfo 对象方法
		.IsDir 方法判断是否为文件夹
		.Name 获取文件名（不含路径）
		.ModTime 为最后修改日期
		.Size 为大小（字节）
	*/
	if info.IsDir() {
		fmt.Printf("Name：%s \nSize:%d\nLastChangeTime：%v\nMode:%v\n", info.Name(), info.Size(), info.ModTime(), info.Mode())
	} else {
		fmt.Printf("Name：%s \nSize：%d \nbytes LastChangeTime: %v\nMode:%v\n", info.Name(), info.Size(), info.ModTime(), info.Mode())
	}

	// geng
}
