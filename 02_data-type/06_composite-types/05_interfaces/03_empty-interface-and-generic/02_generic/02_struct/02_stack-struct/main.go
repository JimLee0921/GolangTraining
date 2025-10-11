package main

import (
	"fmt"
	"strconv"
)

// Stack 模拟栈 可以传入不同类型的元素（每个栈里面的元素类型必须一致）
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func main() {
	/*
		使用泛型实现 stack 等不同数据结构容器
	*/
	// 存放整型数据的栈
	var s1 Stack[int]
	// 存放字符串类型的栈
	var s2 Stack[string]

	for i := 0; i <= 10; i++ {
		s1.Push(i)
		s2.Push(strconv.Itoa(i))
	}

	for {
		if v, ok := s1.Pop(); ok {
			fmt.Println("get value:", v)
		} else {
			fmt.Println("stack s1 is empty")
			break
		}
	}
}
