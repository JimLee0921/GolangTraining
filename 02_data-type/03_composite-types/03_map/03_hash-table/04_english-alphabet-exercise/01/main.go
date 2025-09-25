package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// main 演示如何抓取英语词表原始数据，作为哈希处理的素材来源。
func main() {
	// 下载词表文件并全部读取到内存，简单打印内容验证数据获取成功。
	//res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	//if err != nil {
	//    log.Fatalln(err)
	//}
	//bs, _ := ioutil.ReadAll(res.Body)
	//str := string(bs)
	//fmt.Println(str)

	// 1. 发起请求（上述演示地址已经 403 换成这个 GITHUB 的）
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt")
	// 2. 判断是否出现错误
	if err != nil {
		log.Fatalln(err) // log.Fatalln 会打印错误消息，并且 调用 os.Exit(1) 立即结束程序
	}
	// 3. 读取内容
	body, _ := io.ReadAll(res.Body) // io.ReadAll(一次性把整个响应体读到内存中，返回 []byte)
	// 4. 把[]byte字节转为字符串
	strBody := string(body)
	fmt.Println(strBody) // 少读取一些

}
