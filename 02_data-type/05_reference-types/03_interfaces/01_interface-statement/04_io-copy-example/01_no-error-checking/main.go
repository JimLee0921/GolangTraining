package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	/*
		io.Copy 函数 结合不同实现了 io.Reader 和 io.Writer 接口的对象，实现通用的数据拷贝
		展示了接口的多态性：
			io.Copy 并不知道也不关心给它的 Reader 是字符串、内存还是网络
			它只需要保证有 Read 方法（实现了 io.Reader），就能通用处理。
	*/
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."
	rdr := strings.NewReader(msg) // 拷贝字符串（strings.NewReader）
	io.Copy(os.Stdout, rdr)

	rdr2 := bytes.NewBuffer([]byte(msg)) // 拷贝缓冲区（bytes.NewBuffer）
	io.Copy(os.Stdout, rdr2)

	res, _ := http.Get("https://www.baidu.com") // 拷贝网络响应（http.Get）
	io.Copy(os.Stdout, res.Body)
	res.Body.Close()
}
