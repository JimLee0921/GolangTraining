package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// main 演示使用 bufio.Scanner 拆分文字，说明后续哈希统计可逐词处理
func main() {
	// 构造一段示例文字并用 ScanWords 模式逐词输出，展示扫描流程。
	const input = "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat."
	scanner := bufio.NewScanner(strings.NewReader(input))
	// bufio 内置了几种分割函数：按行 (ScanLines)、按字节 (ScanBytes)、按单词 (ScanWords)。 在这里选择 ScanWords，遇到空格/标点作为分隔符，每次返回一个单词
	scanner.Split(bufio.ScanWords)
	// scanner.Scan() 会不断推进，直到文本读完或出错
	// 每次循环用 scanner.Text() 取出扫描到的单词
	// 逐词打印，可以看到一长段文字被拆成一个个词
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// 检查扫描过程中是否有错误，如果有，就打印到标准错误输出。
	if err := scanner.Err(); err != nil {
		log.Println("reading input:", err)
	}
	// 结论：利用 Scanner 可逐词处理流式文本，为哈希表计数做准备
}
