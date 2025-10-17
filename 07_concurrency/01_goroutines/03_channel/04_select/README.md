## Select

在 Go 并发编程里：

* **channel** 用来在 goroutine 之间传递数据
* **select** 用来同时监听多个 channel 的读写操作，从而实现 **多路复用**（类似网络编程里的 `select`/`poll`/`epoll`）

---

## 基本语法

```go
select {
case val := <-ch1:
fmt.Println("收到 ch1 的数据：", val)
case ch2 <- 10:
fmt.Println("向 ch2 发送数据 10")
default:
fmt.Println("没有可用的 channel 操作")
}
```

解释：

* `case` 分支里可以是 **读操作** 或 **写操作**
* 如果多个 `case` 同时满足，Go 会随机选择一个执行
* 如果都不满足，并且有 `default` 分支，就会执行 `default`，避免阻塞
* 如果没有 `default`，select 会阻塞直到有某个 case 就绪

---

## 使用场景总结

* **同时等待多个通道**：处理最快返回的结果
* **实现超时控制**：配合 `time.After`
* **非阻塞尝试**：配合 `default` 分支
* **优雅退出**：配合 `context.Done()` 来监听取消信号

