package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// main 结合网络请求与 Scanner，将整本书分词并逐个输出，验证数据管线
func main() {
	// 下载《白鲸记》，用 ScanWords 逐词扫描并打印，确认后续可基于该数据做桶统计。
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// 结论：已经具备从真实文本流中逐词获取数据的能力，哈希表可以基于扫描结果统计分布
}
