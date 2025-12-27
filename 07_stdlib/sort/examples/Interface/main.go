package main

import (
	"fmt"
	"sort"
)

// Order 自定义结构体用于定义对应的切片类型
type Order struct {
	SKU   string
	Sales int
}

// BySalesDescThenSku 定义新切片类型实现 sort.Interface 接口
// 不污染原始模型
// 一个排序规则 = 一个类型
// 可以复用、组合、测试
type BySalesDescThenSku []Order

func (o BySalesDescThenSku) Len() int {
	return len(o)
}

func (o BySalesDescThenSku) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o BySalesDescThenSku) Less(i, j int) bool {
	// 先按销量降序排序
	if o[i].Sales != o[j].Sales {
		return o[i].Sales > o[j].Sales
	}

	// 如果销量相同再按 sku 升序排序
	return o[i].SKU < o[j].SKU
}

func main() {
	orders := []Order{
		{SKU: "B100", Sales: 10},
		{SKU: "A200", Sales: 20},
		{SKU: "A100", Sales: 10},
		{SKU: "C300", Sales: 20},
	}
	fmt.Println(orders)
	sort.Sort(BySalesDescThenSku(orders))
	fmt.Println(orders)

	// Reverse 改为销量升序
	sort.Sort(sort.Reverse(BySalesDescThenSku(orders)))
	fmt.Println(orders)
}
