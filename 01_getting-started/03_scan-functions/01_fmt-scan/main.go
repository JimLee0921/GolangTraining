package main

import "fmt"

func main() {
	/*
		标准输入 os.Stdin 读取
		fmt.Scan
			func Scan(a ...any) (n int, err error)
			从标准输入读取，以空格/换行分隔，依次赋值给参数
			返回成功扫描的项数
	*/
	// fmt.Scan
	var name string
	var age int
	fmt.Println("请输入姓名和年龄 (例如: Tom 18):")
	n, err := fmt.Scan(&name, &age)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}
	fmt.Printf("成功读取 %d 个值\n", n)
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)

}
