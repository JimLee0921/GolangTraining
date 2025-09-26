package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// main 在首字母哈希的基础上增加取模，展示如何将大范围编码映射到固定桶数。
func main() {
	// 同样扫描《白鲸记》，但仅使用 12 个桶，通过取模观察更紧凑的分布。
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
	buckets := make([]int, 12) // 十二个桶，通过取模进行分桶
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
	// 结论：通过取模可以把任意编码压缩到有限桶数，为实现哈希表打框架
}

// hashBucket 将单词首字母编码对桶数量取模，返回桶索引。
func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
