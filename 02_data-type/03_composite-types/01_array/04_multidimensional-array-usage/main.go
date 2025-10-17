package main

import "fmt"

// main 演示二维数组的常见用法。
func main() {
	// 1. 定义并初始化二维数组：班级两次考试的成绩表。
	scores := [2][5]int{
		{95, 88, 76, 90, 84},
		{98, 92, 80, 87, 91},
	}
	fmt.Println("scores:", scores)

	// 2. 访问元素：先指定行，再指定列。
	fmt.Println("第一行第三个成绩:", scores[0][2])

	// 3. 获取行数与列数，len 作用于第一维和第二维。
	fmt.Printf("行数: %d，列数: %d\n", len(scores), len(scores[0]))

	// 4. 修改元素：覆盖第二行第四位的成绩。
	scores[1][3] = 93
	fmt.Println("修改后的 scores:", scores)

	// 5. 使用两层 for 循环遍历，配合 break/continue 控制流程。
	for row := 0; row < len(scores); row++ {
		fmt.Printf("第 %d 行: ", row)
		for col := 0; col < len(scores[row]); col++ {
			fmt.Printf("%d ", scores[row][col])
		}
		fmt.Println()
	}

	// 6. for range 遍历，同时获取行索引与每行数据。
	for rowIndex, rowData := range scores {
		fmt.Printf("row %d -> %v\n", rowIndex, rowData)
	}

	// 7. 计算某行或某列的汇总信息。
	total := 0
	for _, score := range scores[0] {
		total += score
	}
	fmt.Printf("第一行总分: %d\n", total)

	// 8. 结合函数操作二维数组，例如传递给求平均分的函数。
	fmt.Printf("全班平均分: %.2f\n", average(scores))
}

// average 计算二维数组所有分数的平均值。
func average(data [2][5]int) float64 {
	sum := 0
	count := 0
	for _, row := range data {
		for _, value := range row {
			sum += value
			count++
		}
	}
	return float64(sum) / float64(count)
}
