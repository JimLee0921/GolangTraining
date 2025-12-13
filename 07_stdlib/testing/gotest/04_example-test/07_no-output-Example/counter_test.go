package counter

// 没有 Output 会被编译，会显示在文档但是不会执行
func ExampleCounter_usage() {
	c := &Counter{1}
	c.Inc()
	c.Inc()
	_ = c
}
