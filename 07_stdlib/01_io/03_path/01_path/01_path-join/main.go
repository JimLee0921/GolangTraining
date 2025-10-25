package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("api", "v1", "user"))
	fmt.Println(path.Join("/a/", "/b/", "c"))
	fmt.Println(path.Join("a", "../b", "c")) // 自动 Clean 操作 去掉冗余路径符号，如 .、..、重复斜杠

	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))

	fmt.Println(path.Join("a/b", "../../../xyz"))

	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))
	/*
		api/v1/user
		/a/b/c
		b/c
		a/b/c
		a/b/c
		a/b/c
		../xyz

		a
		a
	*/
}
