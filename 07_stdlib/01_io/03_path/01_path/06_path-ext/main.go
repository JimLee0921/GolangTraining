package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println("one:", path.Ext("/a/b/c/bar.css"))
	fmt.Println("two:", path.Ext("/a/b"))
	fmt.Println("three:", path.Ext("a.txt"))
	fmt.Println("four:", path.Ext(""))
	/*
		one: .css
		two:
		three: .txt
		four:
	*/
}
