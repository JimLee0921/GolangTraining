# 开始一个 Go 项目

## 初始化模块

```bash
go mod init example/hello
```

上面命令会生成一个 go.mod。`example/hello` 是 `module-path`就是当前项目的模块名，可以设置为成自己的 Github 本项目地址

```text
module module-path

go 1.25
```

## hello world

创建 Go 源文件

```bash
vim hello.go
```

编写第一个 Go 程序

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello World!")
}
```

## 运行

1. 使用 run：`go run hello.go` 可以直接运行 Go 的源码
2. 使用 build：`go build hello.go` 会创建出一个可执行文件 `hello`，直接 `./hello` 即可