# `select` 多路复用（多通道监听）

当在并发程序中有多个通道时：

* 某些 goroutine 可能在发消息
* 某些通道可能在等待
* 有的可能已经超时
* 这是可能希望哪个通道先就绪就先处理

> 这时就需要 `select` —— 可以同时监听多个 channel 的发送与接收操作

---

## 基本语法结构

```
select {
    case v := <-ch1:
    // 从 ch1 接收数据
    case ch2 <- x:
    // 向 ch2 发送数据
    case <-time.After(1 * time.Second):
    // 超时处理
    default:
    // 当所有通道都未准备好时执行
}
```

特点：

* `select` 会随机选择一个已经“准备好”的 case 执行
* 如果没有 case 可执行，且没有 `default`，`select` 会阻塞等待
* 如果有 `default`，则立即执行 `default` 分支（非阻塞）
* 如果有 `time.After`，在规定时间内没有结果会走超时 case

---

## 多通道监听机制

### select + case

`select` 会等待任意一个通道有数据后立即执行对应 case，当两个都准备好时，Go 会随机选一个执行（公平调度）

```
// 使用 select 等待两个通道的数据
for i := 0; i < 2; i++ {
    select {
    case msg1 := <-ch1:
        fmt.Println("receive:", msg1)
    case msg2 := <-ch2:
        fmt.Println("receive:", msg2)
    }
}
```

### 超时机制

`time.After()` 会在指定时间后向返回的通道发送一个值，可以用来实现超时控制或请求取消机制。

```
select {
    case v := <-ch:
    fmt.Println("收到：", v)
    case <-time.After(2 * time.Second):
    fmt.Println("超时！")
}
```

### `default`分支

非阻塞操作， 如果 case 中有 default 关键字，没有准备好的通道不会阻塞，立即执行 `default` 分支，常用于轮询、限流、状态检测

```
select {
    case ch <- 100:
    fmt.Println("成功发送")
    default:
    fmt.Println("通道未准备好，跳过发送")
}
```

### 循环读取多个通道

通道关闭安全循环模式：用 `ch = nil` 避免已关闭通道被再次选中（因为关闭通道总是可读的）

```
for {
    select {
    case msg, ok := <-ch1:
        if !ok {
            ch1 = nil // 防止再次触发此通道
            fmt.Println("ch1 已关闭")
        } else {
            fmt.Println("收到", msg)
        }
    case msg, ok := <-ch2:
        if !ok {
            ch2 = nil
            fmt.Println("ch2 已关闭")
        } else {
            fmt.Println("收到", msg)
        }
    default:
        // 所有通道都暂时没数据
        if ch1 == nil && ch2 == nil {
            fmt.Println("所有通道都已关闭，退出。")
            return
        }
        time.Sleep(100 * time.Millisecond)
    }
}
```

`msg, ok := <-ch`：Go 接收通道的安全写法。

- 如果通道还开着且有数据：ok == true，msg 是收到的值
- 如果通道已关闭且缓存读完：ok == false，msg 是通道元素类型的零值
- `ch1 = nil`
    - 当一个通道被关闭后，它仍然是可读的，只是读到的值都是零值
    - 如果不把它设为 nil，select 会一直选中这个 case，导致程序疯狂地执行 ch 已关闭 分支
    - 当把 ch 置为 nil：它变成一个永远不会被触发的通道；也就是说，select 不会再选中这个 case
    - 等价于把这个 case 从 select 中移除
- `default`分支
    - 当两个通道都暂时没数据时，会走 default
    - 若两个通道都关闭（都变成 nil），则退出循环
    - 否则 Sleep 一下再继续循环

### select + done 通道（优雅退出）

```
done := make(chan struct{})

go func () {
    for {
        select {
            case <-done:
                fmt.Println("收到退出信号")
                return
            default:
                fmt.Println("执行中...")
                time.Sleep(300 * time.Millisecond)
        }
    }
}()

time.Sleep(time.Second)
done <- struct{}{} // 通知退出
time.Sleep(500 * time.Millisecond)
```

输出：

```
执行中...
执行中...
执行中...
收到退出信号
```

> `done` 通道是 Go 并发中最常见的停止信号手段，这里使用空结构体作为信号
>
> `select + context` 发送中途退出信号见 select-and-context

### 行为总结

| 情况                    | 行为                       |
|-----------------------|--------------------------|
| 至少有一个 case 就绪         | 随机选择一个执行                 |
| 所有 case 都阻塞，无 default | `select` 自己阻塞            |
| 所有 case 都阻塞，有 default | 立即执行 default             |
| 通道关闭后                 | 读操作返回零值 + ok=false，依然可选中 |
| 可在 for 中循环使用          | 动态监听多个通道直到全部关闭           |

---

## 与 switch 对比

| 对比项      | `switch` | `select` |
|----------|----------|----------|
| 匹配对象     | 任意表达式    | 通道通信     |
| 是否并发安全   | 无关       | 专为通道设计   |
| 阻塞行为     | 无阻塞      | 可阻塞等待    |
| 多个分支同时就绪 | 只匹配一个    | 随机选一个执行  |

