package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// main 使用 map 嵌套 map 构建桶结构，可直接统计单词出现次数
func main() {
	// 下载《福尔摩斯探案集》，为每个桶维护 map[string]int 计数，统计相同单词频率。
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
	buckets := make(map[int]map[string]int) // 外层：map[int]...，键是桶号（0–11），值是一个内部 map，内层：map[string]int，键是单词，值是出现次数
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int) // 初始化时先给 12 个桶都分配一个空 map
	}
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12) // 算桶号
		buckets[n][word]++        // 单词出现次数+1
	}
	for k, v := range buckets[6] { // 打印第 6 号桶的内容
		fmt.Println(v, " \t- ", k) // 次数 + 单词
	}
	// 结论：借助 map 桶不仅可以分组，还能直接得到词频统计，体现哈希表用途
}

// hashBucket 累加字符编码并取模，得到桶索引
func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
