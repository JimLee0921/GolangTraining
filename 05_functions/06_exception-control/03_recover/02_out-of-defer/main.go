package main

func main() {
	recover() // 无效，程序仍然崩溃

	defer func() {
		recover() // 有效
	}()

	panic("boom")
}
