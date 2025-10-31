package main

import (
	"fmt"
	"io"
	"os"
)

// 文件内容：ABCDEFGHIJKLMNOPQRSTUVWXYZ
func main() {
	// os.File 自动实现了 io.Seeker，所以可以随意移动指针
	// 打开文件，支持读写
	f, err := os.OpenFile("temp_files/test.txt", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 1. 从开头读取前 5 个字符
	f.Seek(0, io.SeekStart) // = io.SeekStart : 从文件开头
	buf := make([]byte, 5)
	f.Read(buf)
	fmt.Println("前 5 个字符:", string(buf)) // ABCDE

	// 2. 跳到第 10 个字节读取 5 个字符
	f.Seek(10, io.SeekStart)
	buf = make([]byte, 5)
	f.Read(buf)
	fmt.Println("从第 10 字节起的 5 个字符:", string(buf)) // KLMNO

	// 3. 从当前再往前回退 3 个字节，并读取 4 个字符
	f.Seek(-3, io.SeekCurrent) // 相对当前位置向前移动 3 字节
	buf = make([]byte, 4)
	f.Read(buf)
	fmt.Println("回退 + 读取:", string(buf)) // MNO

	// 4. 从文件末尾倒数第 5 个字符读取 5 个字符
	f.Seek(-5, io.SeekEnd)
	buf = make([]byte, 5)
	f.Read(buf)
	fmt.Println("末尾倒数 5 字符:", string(buf)) // VWXYZ

	// 5. 在文件第 5 个字符位置写入字符串 (覆盖写)
	f.Seek(5, io.SeekStart)
	f.Write([]byte("----")) // 将 F G H I 覆盖掉

	// 6. 再将文件指针回到开头并输出完整内容
	f.Seek(0, io.SeekStart)
	all, _ := io.ReadAll(f)
	fmt.Println("\n修改后的文件内容:")
	fmt.Println(string(all))

}
