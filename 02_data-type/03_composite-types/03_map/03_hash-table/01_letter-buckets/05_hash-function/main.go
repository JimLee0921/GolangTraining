package main

import "fmt"

// main 展示最简单的哈希函数：取单词首字母后对桶数量取模
func main() {
	// 调用 hashBucket 计算 "Go" 落入的桶位，说明哈希函数输出的是桶索引。
	n := hashBucket("Go", 12)
	fmt.Println(n)
	// 结论：哈希函数负责把任意键映射为固定范围的桶编号。
}

// 把取余分桶的思路真正封装成一个哈希函数 hashBucket
// 以单词首字母编码对桶数取模，返回桶索引
func hashBucket(word string, buckets int) int {
	letter := int(word[0]) // 取字符串的第一个字节，先取一个 byte(unit8，只能表示0-255之间整数太小容易溢出)，再升格为 int(int32)
	fmt.Println(letter)
	bucket := letter % buckets // 取模把任意整数映射到 0 ~ buckets-1
	return bucket
}
