package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// main 使用切片组合构建桶，保存落在同一桶的单词列表
func main() {
	// 下载《福尔摩斯探案集》，根据 hashBucket 结果把单词追加到对应的桶切片中
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
	buckets := make([][]string, 12) // 双层切片，内层切片保存单词，但是因为是切片并不会去重，所有重复的单词也都在一个桶中
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	for i := 0; i < 12; i++ {
		// 打印每个内层切片桶里单词的数量
		fmt.Println(i, " - ", len(buckets[i]))
		fmt.Println(buckets[i])
	}
	// 最后再打印外层切片的长度和容量
	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))
	// 结论：不仅仅用桶去计数，而是让每个桶保存落在这个桶里的完整单词列表，这样可以处理冲突，也能做后续分
}

// hashBucket 累加字符编码后取模，得到桶索引。
func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word { // 累加所有字符码点
		sum += int(v)
	}
	return sum % buckets
	// 如需观察碰撞，可切换为 len(word) % buckets 获得更偏斜的分布
}

/*
AI修改：原版本会 append 空切片初始化，这里改为直接访问固定长度的桶切片，
避免重复扩容；make 创建的外层切片可以直接搭配 append 逐步填充内部数据。
*/
