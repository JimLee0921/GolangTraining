package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// main 演示如何通过 HTTP 获取文本，为哈希统计准备真实的词源数据
func main() {
	// 作者写法
	//// 下载《白鲸记》文本并读取全部字节，输出内容验证网络数据可用。
	//res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bs, err := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s", bs)
	//// 结论：哈希示例后续可以基于真实文本进行分桶与统计。

	// 新版本写法
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalf("http get failed: %v", err)
	}
	// 保证关闭连接
	defer func() {
		if cerr := res.Body.Close(); cerr != nil {
			log.Printf("failed to close response body: %v", cerr)
		}
	}()
	// 检查状态码是否正常
	if res.StatusCode != http.StatusOK {
		log.Fatalf("bad status: %s", res.Status)
	}

	// 读取所有内容
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("read response body failed: %v", err)
	}

	// 打印前 500 个字符，避免控制台刷屏
	fmt.Printf("%s\n", bs[:500])
	fmt.Println("... (truncated)")
} // 结论：哈希示例后续可以基于真实文本进行分桶与统计
