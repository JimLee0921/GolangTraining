package main

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := make(Set[T])
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func main() {
	s := NewSet("A", "B", "C")
	s.Add("D")
	fmt.Println(s.Has("A")) // true
	fmt.Println(s.Has("Z")) // false
}
