# 引用外部包代码

> 具体代码见：[引用外部包](../../01_getting-started/01_tutorial/02_call-external-package/main.go)

## 调用外部模块

* 通过 import "rsc.io/quote" 声明要用外部模块。
* 直接调用 quote.Go() 函数。

```go
package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go()) // 调用外部模块里的函数
}
```

## 下载依赖

* 下载 rsc.io/quote 及其依赖（例如 rsc.io/sampler、golang.org/x/text）
* 把依赖和版本写入 go.mod
* 把完整性校验写入 go.sum

```bash
go mod tidy
```

## 运行

```bash
go mod run .
```

> 输入如下：Don't communicate by sharing memory, share memory by communicating.
