package main

import "fmt"

// main 函数参数和传参时直接使用切片
func main() {
	/*
		函数的arg和param直接使用切片进行传入
	*/
	data := []float64{523, 23, 54, 12, 5.3, 343, 21}
	averageNum := average(data)
	fmt.Println(averageNum)
}

func average(floatSlice []float64) float64 {
	total := 0.0
	for _, value := range floatSlice {
		total += value
	}
	return total / float64(len(floatSlice))
}
