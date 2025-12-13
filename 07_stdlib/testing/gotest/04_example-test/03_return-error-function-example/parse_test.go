package parse

import "fmt"

// Parse 会返回 error 类型，这里 nil 得写成 <nil>，并且多个返回测试值使用空格分开
func ExampleParse() {
	v1, err1 := Parse("123")
	v2, err2 := Parse("234")
	fmt.Println(v1, err1)
	fmt.Println(v2, err2)
	// Output:
	// 123 <nil>
	// 234 <nil>
}
