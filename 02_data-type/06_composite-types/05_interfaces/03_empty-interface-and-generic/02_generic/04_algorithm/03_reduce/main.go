package main

import "fmt"

func Reduce[T any, R any](src []T, init R, f func(R, T) R) R {
	acc := init
	for _, v := range src {
		acc = f(acc, v)
	}
	return acc
}

func main() {
	/*
		聚合计算
			泛型 Reduce 是函数式聚合的核心
			类型参数 T（元素类型）与 R（累积结果类型）可以不同
			灵活、安全、零开销
	*/
	nums := []int{1, 2, 3, 4}
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println(sum) // 10
}
