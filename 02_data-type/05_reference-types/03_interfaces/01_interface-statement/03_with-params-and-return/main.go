package main

import (
	"fmt"
)

// Payer 定义接口
type Payer interface {
	Pay(amount float64) string
	Refund(amount float64) bool
}

// CreditCard 实现1：信用卡支付
type CreditCard struct {
	Holder string
	Limit  float64
}

func (c *CreditCard) Pay(amount float64) string {
	if amount > c.Limit {
		return fmt.Sprintf("%s 的信用卡余额不足（限额 %.2f）", c.Holder, c.Limit)
	}
	c.Limit -= amount
	return fmt.Sprintf("%s 用信用卡支付了 %.2f 元，剩余额度 %.2f", c.Holder, amount, c.Limit)
}

func (c *CreditCard) Refund(amount float64) bool {
	c.Limit += amount
	fmt.Printf("退回 %.2f 元到 %s 的信用卡，现在额度 %.2f\n", amount, c.Holder, c.Limit)
	return true
}

// Cash 实现2：现金支付
type Cash struct {
	Balance float64
}

func (c *Cash) Pay(amount float64) string {
	if amount > c.Balance {
		return fmt.Sprintf("现金不足，余额 %.2f", c.Balance)
	}
	c.Balance -= amount
	return fmt.Sprintf("成功用现金支付 %.2f 元，剩余余额 %.2f", amount, c.Balance)
}

func (c *Cash) Refund(amount float64) bool {
	c.Balance += amount
	fmt.Printf("退回 %.2f 元现金，现在余额 %.2f\n", amount, c.Balance)
	return true
}

// Wallet 实现3：数字钱包
type Wallet struct {
	User   string
	Amount float64
}

func (w *Wallet) Pay(amount float64) string {
	if amount > w.Amount {
		return fmt.Sprintf("%s 的钱包余额不足", w.User)
	}
	w.Amount -= amount
	return fmt.Sprintf("%s 用钱包支付 %.2f 元，剩余 %.2f", w.User, amount, w.Amount)
}

func (w *Wallet) Refund(amount float64) bool {
	w.Amount += amount
	fmt.Printf("退回 %.2f 元到 %s 的钱包，现在余额 %.2f\n", amount, w.User, w.Amount)
	return true
}

// ProcessPayment 公共函数：执行支付操作
func ProcessPayment(p Payer, amount float64) {
	fmt.Println(p.Pay(amount))
}

// ProcessRefund 公共函数：执行退款操作
func ProcessRefund(p Payer, amount float64) {
	p.Refund(amount)
}

func main() {
	c1 := &CreditCard{"Alice", 500}
	cash := &Cash{200}
	w := &Wallet{"Bob", 20}

	// 支付操作
	ProcessPayment(c1, 150)
	ProcessPayment(cash, 80)
	ProcessPayment(w, 50)

	// 退款操作
	ProcessRefund(c1, 10)
	ProcessRefund(cash, 20)
	ProcessRefund(w, 40)

}
