package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// main 基于首字母哈希统计单词数量，展示最朴素的分桶计数方法
func main() {
	// 下载《白鲸记》，为每个单词计算 hashBucket，并在对应切片桶中自增次数。
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer func() {
		if cerr := res.Body.Close(); cerr != nil {
			log.Printf("failed to close response body: %v", cerr)
		}
	}()
	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 200) // make([]int, 200) 创建了一个长度为 200 的切片（元素初始全是 0），充当桶数组
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}
	// 二十四个英文字母的整数值只需要看 64 - 123 即可
	for i := 65; i < 123; i++ {
		fmt.Printf("%c : %d\n", i, buckets[i])
	}
	// 结论：即便哈希函数很简单，也能统计不同首字母落桶的次数分布
}

// hashBucket 返回单词首字节的整数值，作为桶索引
func hashBucket(word string) int {
	return int(word[0])
}
