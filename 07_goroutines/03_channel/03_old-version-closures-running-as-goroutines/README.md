这里演示的是 Go 1.22 之前的版本中在并发中使用闭包可能会引起一些混淆

文档见：https://go.dev/doc/faq#closures_and_goroutines

```go
package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// 等待所有 goroutines 完成后再退出
	for _ = range values {
		<-done
	}
}

```

这里 1.22 之前的版本输出的不是 a b c ，而可能是 c c c

这是因为循环的每次迭代都使用变量 的同一个实例，所以每个闭包共享该变量。闭包运行时，它会打印执行时v的值，但自 goroutine
启动以来可能已被修改。

为了在启动每个闭包时将其当前值绑定v到闭包，必须修改内循环，使其在每次迭代时创建一个新变量。一种方法是将变量作为参数传递给闭包：

```go
 for _, v := range values {
go func ( u string) {
fmt.Println( u )
done <- true 
}( v )
}

```

更简单的方法是创建一个新变量：

```go

for _, v := range values {
v := v // 创建一个新的 'v'
go func () {
fmt.Println( v )
done <- true 
}()
}
```