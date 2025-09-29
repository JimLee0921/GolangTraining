package main

import "fmt"

func main() {
	/*
		输出从 1 到 N 的数字（一般是 100），但有规则：
			如果数字能被 3 整除 -> 输出 "Fizz"
			如果数字能被 5 整除 -> 输出 "Buzz"
			如果同时能被 3 和 5 整除 -> 输出 "FizzBuzz"
			其他情况 -> 输出数字本身
	*/
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
