package main

import (
	"fmt"
	"math"
)

func main() {
	// ----- math -----
	fmt.Println(math.Abs(-3.5))
	fmt.Println(math.Max(3, 5), math.Min(3, 5))
	fmt.Println(math.Ceil(2.3), math.Floor(2.8), math.Round(2.5), math.Trunc(-2.9))

	fmt.Println(math.Pow(2, 3), math.Sqrt(9), math.Log(math.E), math.Log10(100))

	// 三角函数（注意：弧度制）
	rad := 30 * math.Pi / 180
	fmt.Println(math.Sin(rad), math.Cos(rad), math.Tan(rad))

	// 向量长度（常用计算距离/长度）
	fmt.Println(math.Hypot(3, 4)) // 5
}
