## goroutine 的同步与等待

goroutine 是异步执行的：

```
go worker()
fmt.Println("main done")
```

上面主函数（main goroutine）不会等待 `worker()` 结束就退出。
而一旦主 goroutine 结束，整个程序都会终止。

因此，为了确保 goroutine 都执行完再退出程序，需要一种同步机制来等待它们的完成。

最常用的就是：`sync.WaitGroup`

---

### `sync.WaitGroup` 概述

`WaitGroup` 是 Go 提供的计数器同步工具。
它的核心思想：

* 用一个计数器记录还有多少 goroutine 没做完
* 当所有 goroutine 完成时，主 goroutine 才继续执行

---

### WaitGroup 基本使用

| 方法               | 作用                                  |
|------------------|-------------------------------------|
| `Add(delta int)` | 增加（或减少）等待的 goroutine 数              |
| `Done()`         | 当前 goroutine 完成任务后调用（相当于 `Add(-1)`) |
| `Wait()`         | 阻塞等待，直到计数器变为 0                      |

```
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // 任务完成后计数 -1
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)          // 计数 +1
        go worker(i, &wg)  // 启动 goroutine
    }

    wg.Wait() // 阻塞等待所有 worker 完成
    fmt.Println("All workers finished.")
}
```

---

### 工作原理图

```
+-------------------+
| 主 goroutine (main)|
+---------+---------+
          |
+----------v-----------+
|  wg.Add(1) 启动 goroutine |
+------------------------+
          |
+-------------------+
|  每个 worker 执行  |
|  defer wg.Done()  |
+-------------------+
          |
 （当 wg 计数归零）
          |
+-------------------+
|     wg.Wait() 返回 |
+-------------------+
```

---

### 关键点与注意事项

#### `wg.Add()` 必须在启动 goroutine 之前调用

```
go worker(&wg)
wg.Add(1)
```

结果是：

* goroutine 太快执行到 `wg.Done()`
* 计数器变为负数 -> panic: `sync: negative WaitGroup counter`

正确做法：

```
wg.Add(1)
go worker(&wg)
```

---

#### 必须用 指针传递 `*sync.WaitGroup`

因为多个 goroutine 要共享同一个计数器。

```
func worker(wg *sync.WaitGroup) { ... } // 正确
```

如果用值传递（非指针），每个 goroutine 都会拷贝一个副本 -> 没有效果。

---

#### Done()` 通常搭配 `defer`

让函数不管正常还是异常退出，都能安全地调用：

```
defer wg.Done()
```

---

### 对比 time.Sleep

| 方式               | 特点              | 场景         |
|------------------|-----------------|------------|
| `time.Sleep()`   | 简单粗暴，无法精确控制等待时机 | 调试、临时测试    |
| `sync.WaitGroup` | 等待确切任务完成        | 正式并发逻辑（推荐） |

### 最大进程并行度

可以在 main 包的 init 函数中进行定义：

```
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
```

runtime.GOMAXPROCS(runtime.NumCPU())

- 使用尽可能多的 CPU 核心来同时运行 goroutines（这里就是把逻辑 CPU 核心数设为最大值）
- 默认情况下 Go 可能不会使用所有 CPU，把它调到 NumCPU() 后，可以充分并行
- GOMAXPROCS 设置后，如果机器有多核，两个 goroutine 可能真的是并行在不同核上执行，而不是单核抢时间片



