package main

import "fmt"

// main switch 语句基础写法
func main() {
	/*
		switch 语句 分支可以有多个，
		default 是可以省略的如果省略 default 且任何分支都没有匹配到则什么都不做
		表达式也是可以省略的，省略条件相当于 switch true
		写条件判断，就要用 不带表达式的 switch，也就是 switch {}，这种情况每个 case 必须是布尔表达式
		单个 case 还支持多个值使用 `,` 分割进行匹配
			switch 条件表达式 {
		    case 1:
		        匹配 case1 代码体
		    case 2:
		        匹配 case2 代码体
		    case 3:
		        匹配 case3 代码体
		    default:
		        什么都没匹配到代码体
		    }
	*/
	name := "JimLee"
	switch name {
	case "Rose":

		fmt.Println("I'm Rose")
	case "Jack":

		fmt.Println("I'm Jack")
	case "JimLee":

		fmt.Println("I'm JimLee")
	case "Doll":
		fmt.Println("I'm Doll")
	default:
		fmt.Println("I'm Nobody")
	}

	// 省略 条件，每个 case 后必须是 bool 表达式
	age := 55
	switch {
	case age < 18:
		fmt.Println("未成年")
	case age < 30:

		fmt.Println("成年")
	case age < 40:
		fmt.Println("壮年")
	case age < 50:
		fmt.Println("壮年")
	default:
		fmt.Println("老年")
	}

	// 省略 default 未匹配到就什么都不做
	x := 4
	switch x {
	case 1:
		fmt.Println("x=1")
	case 2:
		fmt.Println("x=2")
	case 3:
		fmt.Println("x=3")
	case 4:
		fmt.Println("x=4")
	case 5:
		fmt.Println("x=5")
	}

	// 单个 case 设置多个值
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("周末")
	default:
		fmt.Println("工作日")
	}

}
