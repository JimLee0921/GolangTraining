package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

func main() {
	/*
		泛型容器写一次，支持任意类型
		不需要 interface{} + 断言
		编译期生成 Stack[string], Stack[int] 等版本
	*/
	var s Stack[string]
	s.Push("A")
	s.Push("B")

	fmt.Println(s.Pop()) // B true
	fmt.Println(s.Pop()) // A true
}
