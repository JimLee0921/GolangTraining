package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReplaceDemo() {
	r := strings.NewReplacer(
		"foo", "FOO",
		"bar", "BAR",
		"baz", "BAZ",
	)

	input := "foo and bar but not foobar or baz"
	output := r.Replace(input)
	fmt.Println(output)
}

func ShowConReplace() {
	r := strings.NewReplacer("a", "b", "b", "c")

	fmt.Println(r.Replace("a"))  // 输出的是 "b"
	fmt.Println(r.Replace("b"))  // 输出的是 "c"
	fmt.Println(r.Replace("ab")) // 输出的是 "bc"
}

func WriteStringDemo() {
	r := strings.NewReplacer(
		"<", "&lt;",
		">", "&gt;",
		"&", "&amp;",
	)
	input := "<div>Hello & Welcome</div>"

	// 直接将结果写到 stdout，也可以换成文件等
	n, err := r.WriteString(os.Stdout, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nwrite", n, "bytes")
}

func main() {
	ReplaceDemo()
	ShowConReplace()
	WriteStringDemo()
}
