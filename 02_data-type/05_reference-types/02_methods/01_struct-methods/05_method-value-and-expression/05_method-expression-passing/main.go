package main

import "fmt"

type Account struct {
	Balance int
}

// Deposit 存钱方法
func (a *Account) Deposit(n int) {
	a.Balance += n
}

// Repeat 高阶函数：重复调用某个函数
func Repeat(f func(*Account, int), a *Account, n int, times int) {
	for i := 0; i < times; i++ {
		f(a, n)
	}
}

func main() {
	acc := Account{Balance: 10}

	// 创建方法表达式，不绑定对象
	f := (*Account).Deposit
	// 当作普通函数进行传递(Repeat 是普通函数而不是方法，acc 需要手动取地址)
	Repeat(f, &acc, 10, 3) // 调用三次repeat，每次存入10块钱
	fmt.Println("Final Balance", acc.Balance)

}
