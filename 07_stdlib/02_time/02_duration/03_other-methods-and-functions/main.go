package main

import (
	"fmt"
	"time"
)

func CompareDemo() {
	// Duration 比较
	d1 := 3 * time.Second
	d2 := 5 * time.Second
	fmt.Println(d1 < d2)  // true
	fmt.Println(d1 > d2)  // false
	fmt.Println(d1 == d2) // true
}

func NegativeDemo() {
	// 取负
	now := time.Now()
	past := now.Add(-30 * time.Minute)
	fmt.Println(past)
}

func RoundAndTruncateDemo() {
	// 四舍五入和向下取整
	d := 1499 * time.Millisecond
	fmt.Println(d.Round(time.Second))    // 1s
	fmt.Println(d.Truncate(time.Second)) // 1s

	d = 1501 * time.Millisecond
	fmt.Println(d.Round(time.Second))    // 2s
	fmt.Println(d.Truncate(time.Second)) // 1s
}

func DivideDemo() {
	// 相除
	d := 125 * time.Second
	minutes := d / time.Minute // 2
	seconds := d % time.Minute / time.Second
	fmt.Println(minutes, seconds)
}

func AbsDemo() {
	// 绝对值
	positiveDuration := 5 * time.Second
	negativeDuration := -3 * time.Second

	absPositive := positiveDuration.Abs()
	absNegative := negativeDuration.Abs()

	fmt.Printf("Absolute value of positive duration: %v\n", absPositive)
	fmt.Printf("Absolute value of negative duration: %v\n", absNegative)
}

func SinceAndUntilDemo() {
	// Since 是 time.Now().Sub(t) 的简写
	// Until 是 t.Sub(time.Now()) 的简写
	futureTime := time.Now().Add(10 * time.Second)
	durationUntil := time.Until(futureTime)
	fmt.Println(durationUntil)          // 10s
	fmt.Println(time.Since(futureTime)) // -10s

}

func main() {
	CompareDemo()
	NegativeDemo()
	RoundAndTruncateDemo()
	DivideDemo()
	AbsDemo()
	SinceAndUntilDemo()

}
