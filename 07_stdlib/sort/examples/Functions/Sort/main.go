package main

import (
	"fmt"
	"sort"
)

type Order struct {
	SKU   string
	Sales int
}

type BySalesDesc []Order

func (o BySalesDesc) Len() int           { return len(o) }
func (o BySalesDesc) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o BySalesDesc) Less(i, j int) bool { return o[i].Sales > o[j].Sales } // 降序

func main() {
	orders := []Order{{"A", 10}, {"B", 10}, {"C", 20}}
	sort.Sort(BySalesDesc(orders))
	fmt.Println(orders)
}
