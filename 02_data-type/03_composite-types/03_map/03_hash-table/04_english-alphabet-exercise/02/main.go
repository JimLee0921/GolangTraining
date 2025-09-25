package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// main 将词表读入 map 以做去重，展示哈希表在字典数据处理中的应用
func main() {
	// 下载词表，使用 Scanner 按词读取，并存入 map[string]string 达到快速查重的目的
	//res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	//if err != nil {
	//    log.Fatalln(err)
	//}
	//words := make(map[string]string)
	//sc := bufio.NewScanner(res.Body)
	//sc.Split(bufio.ScanWords)
	//for sc.Scan() {
	//    words[sc.Text()] = ""
	//}
	//if err := sc.Err(); err != nil {
	//    fmt.Fprintln(os.Stderr, "reading input:", err)
	//}
	//i := 0
	//for k := range words {
	//    fmt.Println(k)
	//    if i == 200 {
	//        break
	//    }
	//    i++
	//}
	// 下载词表，使用 Scanner 按词读取，并存入 map[string]string 达到快速查重的目的
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt")
	if err != nil {
		log.Fatalln(err)
	}
	words := make(map[string]string)
	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// 使用 map 自动去重
		words[scanner.Text()] = ""
	}
	// 控制打印个数
	i := 0
	for key, _ := range words {
		// map 是无序的,每次打印结果可能都不一样
		fmt.Println(key)
		if i == 200 {
			break
		}
		i++
	}
	// 结论：map 结构天生适合字典数据的去重和快速查询，这正是哈希表的核心价值
}
