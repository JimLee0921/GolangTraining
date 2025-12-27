package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := "HEAD:你好,world!\nBODY:abc123\n"

	// 1. 构造 Reader：把 string 变成“可读流”
	r := strings.NewReader(s)
	fmt.Printf("[init] Size=%d Len=%d\n", r.Size(), r.Len())

	// 2. Read：按字节批量顺序读取，会推进读指针
	buf := make([]byte, 5)
	n, err := r.Read(buf)
	fmt.Printf("[Read] n=%d err=%v data=%q  (Len now=%d)\n", n, err, buf[:n], r.Len())

	// 3. Seek：跳回开头（像文件一样）
	pos, err := r.Seek(0, io.SeekStart) // 从开头算，移动0个字节的位置
	fmt.Printf("[SeekStart] pos=%d err=%v (Len now=%d)\n", pos, err, r.Len())

	// 4. ReadRune + UnreadRune：按 UTF-8 解码读取一个字符，并回退一次
	ch, size, err := r.ReadRune()
	fmt.Printf("[ReadRune] ch=%q size=%d err=%v (Len now=%d)\n", ch, size, err, r.Len())

	err = r.UnreadRune()
	fmt.Printf("[UnreadRune] err=%v (Len now=%d)\n", err, r.Len())

	// 再读一次 rune（应与上一次相同）
	ch, size, err = r.ReadRune()
	fmt.Printf("[ReadRune again] ch=%q size=%d err=%v (Len now=%d)\n", ch, size, err, r.Len())

	// 5. ReadByte：读一个字节（这里读取 ':' 之前的内容会逐步推进）
	// 先 Seek 到 "HEAD" 后面的位置：偏移 4（"HEAD" 是 4 个字节）
	_, _ = r.Seek(4, io.SeekStart)
	b, err := r.ReadByte()
	fmt.Printf("[ReadByte] byte=%q err=%v (pos=%d Len now=%d)\n", b, err, posOf(r), r.Len())

	// 6. ReadAt：随机读取，不改变当前读指针
	peek := make([]byte, 4)
	n, err = r.ReadAt(peek, 0) // 从开头读 4 个字节，应为 "HEAD"
	fmt.Printf("[ReadAt] n=%d err=%v data=%q  (Len unchanged=%d)\n", n, err, peek[:n], r.Len())

	// 7. WriteTo：把“剩余未读”的内容直接写到某个 io.Writer（这里写到 stdout）
	fmt.Println("\n[WriteTo stdout] remaining bytes below:")
	wrote, err := r.WriteTo(os.Stdout)
	fmt.Printf("\n[WriteTo] wrote=%d err=%v (Len now=%d)\n", wrote, err, r.Len())

	// 8. Reset：复用同一个 Reader，换一个新的 string 从头读
	r.Reset("NEW:再来一段\n")
	fmt.Printf("\n[Reset] Size=%d Len=%d\n", r.Size(), r.Len())
	all, _ := io.ReadAll(r)
	fmt.Printf("[ReadAll after Reset] %q\n", string(all))
}

// posOf 只是为了示意当前位置，不属于标准 API；这里通过 Size-Len 计算出来
func posOf(r *strings.Reader) int64 {
	return r.Size() - int64(r.Len())
}
