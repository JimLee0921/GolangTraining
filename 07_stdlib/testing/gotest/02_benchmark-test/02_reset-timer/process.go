package demo

// 模拟一些比较耗时的数据准备工作
func makeHugeData() []int {
	data := make([]int, 1_000_000) // 1_000_000 等于 1000000 下划线用于数字字面量作为分隔符可以增加可读性
	for i := range data {
		data[i] = i
	}
	return data
}

func Process(data []int) int {
	sum := 0
	for i := range data {
		sum += i
	}
	return sum
}
