package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// main 通过改进哈希函数（累加所有字符），比较更均匀的桶分布效果
func main() {
	// 扫描《福尔摩斯探案集》，对每个单词求字符和取模，统计 12 个桶的命中次数
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
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
	buckets := make([]int, 12)
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
	// 结论：合理的哈希函数能让桶分布更均匀，减少碰撞压力
}

// hashBucket 累加单词所有字符的 Unicode 编码，再对桶数取模
func hashBucket(word string, buckets int) int {
	/*
		不再只用首字母，而是把整个单词的所有字符码点累加
		这样每个单词的哈希值差异更大，分布也会更均匀
		最终再取模，映射到 0 ~ buckets-1 的范围
	*/
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
