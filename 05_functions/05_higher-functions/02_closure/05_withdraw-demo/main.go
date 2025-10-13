package main

import "fmt"

func main() {
	/*
		闭包创建的实例互不影响
	*/
	accountOneDeposit, accountOneWithdraw, accountOneGetBalance := bankAccount(500)
	accountTwoDeposit, accountTwoWithdraw, accountTwoGetBalance := bankAccount(500)
	accountOneDeposit(50)   // 账号 one 存入 50
	accountTwoDeposit(8000) // 账号 two 存入 8000
	fmt.Println(accountOneGetBalance())
	accountOneWithdraw(600) // 账号 one 取出 600
	accountTwoWithdraw(500) // 账号 two 取出 500
	fmt.Println(accountTwoGetBalance())
}

/*
定义一个 bankAccount 返回三个闭包函数：存钱、取钱、查余额
*/
func bankAccount(initialBalance int) (deposit func(int), withdraw func(int) bool, getBalance func() int) {
	balance := initialBalance
	deposit = func(amount int) {
		if amount <= 0 {
			fmt.Println("存款金额必须大于 0")
			return
		}
		balance += amount
	}
	withdraw = func(amount int) bool {
		if amount > balance {
			fmt.Println("取款失败，金额不足")
			return false
		}
		balance -= amount
		fmt.Println("取款成功")
		return true
	}
	getBalance = func() int {
		return balance
	}
	return
}
