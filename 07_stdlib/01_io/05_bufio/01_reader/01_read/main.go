package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadString(str string) {
	// strings.NewReader 把普通字符串当作可读取的数据流（实现了 io.Reader 接口的对象）来使用
	reader := bufio.NewReader(strings.NewReader(str))
	buf := make([]byte, 8)

	for {
		n, err := reader.Read(buf)
		// %q 占位符输出 ASCII 编码的字符，相当于 string(buf[:n])
		if err == io.EOF {
			fmt.Println("read ending")
			break
		} else if err != nil {
			fmt.Println("unknown error:", err)
			break
		}
		fmt.Printf("read %d bytes: %q\n", n, buf[:n])
	}
}

func ReadFile(filePath string) {
	f, _ := os.Open(filePath)
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 16)
	for {
		// bufio.Reader 在内部做了缓冲，避免频繁系统调用（相当于一次读多一些数据放在内存里）
		n, err := reader.Read(buf)
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	// 创建 Reader
	ReadString("hello brother shu, wtffffff??????")
	//ReadFile("temp_files/4053322876")
}
