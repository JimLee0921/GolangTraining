package main

func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}

/*
	go test -v ./06_testing/02_unit-test/
	运行结果如下：
		=== RUN   TestFib
		--- PASS: TestFib (0.00s)
		PASS
		ok      github.com/JimLee0921/GolangTraining/06_testing/02_unit-test    0.008s
*/
