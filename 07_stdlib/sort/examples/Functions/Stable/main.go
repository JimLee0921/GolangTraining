package main

import (
	"fmt"
	"sort"
)

type Item struct {
	ID    string
	Group int
}

type ByGroup []Item

func (x ByGroup) Len() int           { return len(x) }
func (x ByGroup) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x ByGroup) Less(i, j int) bool { return x[i].Group < x[j].Group }

func main() {
	items := []Item{
		{ID: "first", Group: 1},
		{ID: "second", Group: 1},
		{ID: "third", Group: 2},
	}
	sort.Stable(ByGroup(items))
	fmt.Println(items) // first、second 的相对顺序会被保留
}
