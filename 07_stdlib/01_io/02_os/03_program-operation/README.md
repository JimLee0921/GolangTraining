# 程序自身相关

Go 的 `os` 包中，有一类函数用于访问当前程序进程的信息，比如：

| 分类          | 函数 / 变量                            | 说明            |
|-------------|------------------------------------|---------------|
| **命令行参数**   | `os.Args`                          | 获取命令行参数（切片）   |
| **可执行文件信息** | `os.Executable()`、`os.Getwd()`     | 获取程序路径、当前工作目录 |
| **退出程序**    | `os.Exit(code)`                    | 立即退出程序        |
| **进程信息**    | `os.Getpid()`、`os.Getppid()`       | 获取当前 / 父进程 ID |
| **用户与主机**   | `os.Hostname()`                    | 获取主机名         |
| **标准输入输出流** | `os.Stdin`、`os.Stdout`、`os.Stderr` | 输入输出控制        |
| **信号与中断**   | 搭配 `os/signal` 包使用                 | 捕获 Ctrl+C 等中断 |

## 获取命令行参数 `os.Args`

`os.Args` 保存命令行参数，以程序名称作为第一个元素

```
var Args []string
```

* 是一个 `[]string` 切片，包含程序启动时传入的所有参数
* **第 0 个元素**（`os.Args[0]`）永远是程序自身的路径
* 后面的元素是实际的命令行参数，只有当程序是通过命令行启动时，才能接收到外部传入的参数
* `go run main.go arg1 arg2 arg3 ...`

---

## 可执行文件与工作目录

### 当前程序的完整路径 `os.Executable()`

语法：`os.Executable() (string, error)`

Executable 返回启动当前进程的可执行文件的路径名。无法保证该路径仍然指向正确的可执行文件。
如果使用符号链接启动进程，则根据操作系统的不同，结果可能是符号链接或其指向的路径。
除非发生错误，否则可执行文件将返回绝对路径。
主要适用于查找相对于可执行文件的资源。

```
path, _ := os.Executable()
fmt.Println("program path:", path)
```

---

### 获取当前工作目录 `os.Getwd()`

`Getwd` 返回当前目录对应的绝对路径名。如果当前目录可以通过多条路径（例如符号链接）到达，`Getwd` 可能会返回其中任意一条路径。
在 Unix 平台上，如果环境变量 PWD 提供了一个绝对名称，并且它是当前目录的名称，则返回它。

```
dir, _ := os.Getwd()
fmt.Println("当前工作目录:", dir)
```

- 相当于 shell 中的 `pwd`
- 如果 `chdir()` 到别的目录，`Getwd()` 也会随之变化

---

### 切换当前工作目录 `os.Chdir(path)`

Chdir 将修改当前工作目录，只能修改为目录。如果发生错误，则返回 `*PathError` 类型错误。

语法：`Chdir() error`

```
os.Chdir("/tmp")
fmt.Println(os.Getwd()) // 输出：/tmp
```

---

## 程序退出

### `os.Exit`

语法：`os.Exit(code int)`

Exit 使当前程序以给定的状态码退出。通常状态码 0 表示成功，非 0 表示错误。程序立即终止；延迟函数 `defer `不会运行。

```
if err != nil {
    fmt.Println("出错，退出程序")
    os.Exit(1)
}
```

> 与 `return` 不同，`os.Exit()` 不会执行 `defer`。

---

## 进程与用户信息

| 类别           | 函数                       | 说明              |
|--------------|--------------------------|-----------------|
| **进程 ID 信息** | `Getpid()` / `Getppid()` | 当前进程 / 父进程 ID   |
| **用户身份信息**   | `Getuid()` / `Geteuid()` | 用户 ID / 有效用户 ID |
| **组身份信息**    | `Getgid()` / `Getegid()` | 组 ID / 有效组 ID   |
| **附属组信息**    | `Getgroups()`            | 所属的所有附属组 ID     |
| **主机名**      | `Hostname()`             | 当前主机名称          |

- `os.Getpid()`：获取当前进程（Process）的 ID
- `os.Getppid()`：获取当前进程的父进程（Parent Process）ID
- `os.Getuid()`：获取当前用户的 真实用户 ID (UID)
    - 在 Linux/Mac 下，这个值对应 /etc/passwd 中的用户 ID
    - 普通用户一般是非零值（如 1000），root 用户为 0
- `os.Geteuid()`：获取当前用户的 有效用户 ID (EUID)
- `os.Getgid()`：获取当前用户所属的真实组 ID (GID)
- `os.Getegid()`：获取当前用户的有效组 ID (EGID)
- `os.Getgroups()`：获取当前用户所属的所有附属组 ID 列表，这些数字对应 /etc/group 文件中列出的组 ID
- `os.Hostname()`：获取当前主机名

### 真实用户和有效用户

| 类型   | 英文全称          | 含义          | 举例         |
|------|---------------|-------------|------------|
| UID  | User ID       | 启动程序的真实用户   | 1000（普通用户） |
| EUID | Effective UID | 当前程序以谁的身份运行 | 0（root）    |
| GID  | Group ID      | 启动用户所属的主组   | 1000       |
| EGID | Effective GID | 当前进程的生效组    | 0（提权后）     |

### 跨平台问题

> Getuid, Geteuid, Getgid, Getegid, Getgroups 在 Windows 上通常返回 0 或空列表
> 这是因为 Windows 没有传统 Unix 的 UID/GID 概念
> 但 Hostname、Getpid、Getppid 是跨平台可用的
---

## 标准输入输出（Stdin / Stdout / Stderr）

`os.Stdin`、`os.Stdout`、`os.Stderr` 是 Go 中与 标准输入输出流（Standard I/O streams） 相关的三个全局变量。
它们是操作系统为每个进程默认分配的 文件流（file descriptors）：

### 文件描述符

操作系统为每个进程维护一个文件表，
每打开一个文件（包括键盘、网络连接、终端等），都会分配一个编号——文件描述符。

| 描述符编号 | 名称           | Go 对应对象     | 默认用途 | 默认绑定     |
|-------|--------------|-------------|------|----------|
| `0`   | 标准输入（stdin）  | `os.Stdin`  | 读入   | 键盘 / 输入源 |
| `1`   | 标准输出（stdout） | `os.Stdout` | 写出   | 终端屏幕     |
| `2`   | 标准错误（stderr） | `os.Stderr` | 写出   | 终端屏幕     |

### 重定向操作符（Redirection Operators）

Shell 提供了几种操作符来改变这些文件描述符的目标（或来源）。

| 操作符    | 含义                          | 默认作用对象                     | 举例                                        |      
|--------|-----------------------------|----------------------------|-------------------------------------------|
| `>`    | 把输出写入文件（覆盖）                 | `stdout (1)`               | `go run main.go > out.txt`                |         
| `>>`   | 把输出追加到文件末尾                  | `stdout (1)`               | `go run main.go >> log.txt`               |               
| `<`    | 从文件读取输入                     | `stdin (0)`                | `go run main.go < input.txt`              |          
| `2>`   | 把错误输出写入文件                   | `stderr (2)`               | `go run main.go 2> err.txt`               |          
| `2>>`  | 把错误追加到文件末尾                  | `stderr (2)`               | `go run main.go 2>> err.log`              |          
| `2>&1` | 把 stderr 重定向到 stdout        | 合并 stdout + stderr         | `go run main.go > all.txt 2>&1`           |          
| `&>`   | 同时重定向 stdout 和 stderr       | bash专用                     | `go run main.go &> all.txt`               |         
| `\|`   | 管道(含错误流)，把一个命令的输出传给下一个命令的输入 | stdout+stderr -> 下一个 stdin | `cat data.txt \| grep "Go"      \| wc -l` |

### Go 中标准流

这些流让的程序可以与用户或外部管道交互：

- 读输入 -> os.Stdin
- 打印输出 -> os.Stdout
- 打印错误信息 -> os.Stderr

这三个变量的类型都是：

```
var Stdin  *File
var Stdout *File
var Stderr *File
```

它们本质上是 *os.File 对象。因此可以像操作文件一样用 Read() / Write() 操作它们。

> `fmt.Fprint` / `fmt.Fprintf` / `fmt.Fprintln`
> 就是 `fmt.Print` / `fmt.Printf` / `fmt.Println` 的“底层通用版本”，
> 区别仅在于：可以自己指定要输出到哪里（比如 io.Writer），而不是固定输出到 os.Stdout。

默认行为

| 函数                                         | 默认输出位置              | 等价写法                         |
|--------------------------------------------|---------------------|------------------------------|
| `fmt.Print` / `fmt.Println` / `fmt.Printf` | **标准输出（os.Stdout）** | `fmt.Fprint(os.Stdout, ...)` |
| `log.Print` / `log.Println`                | **标准错误（os.Stderr）** | `fmt.Fprint(os.Stderr, ...)` |

### `os.Stdin` 标准输入

从用户输入或管道读取数据，使用管道操作见上面

```
fmt.Print("Enter your name: ")
reader := bufio.NewReader(os.Stdin)
name, _ := reader.ReadString('\n')
fmt.Println("Hello,", name)
```

### `os.Stdout` 标准输出

普通的 `fmt.Print` / `fmt.Println` / `fmt.Printf` 默认其实就是往 `os.Stdout` 写，使用管道操作见上面

```
fmt.Fprintln(os.Stdout, "This goes to standard output")
// 等价于
fmt.Println("This goes to standard output")
```

### `os.Stderr` 标准错误

打印错误信息，不与正常输出混在一起，使用管道操作见上面

```
fmt.Fprintln(os.Stderr, "Error: something went wrong")
// 等价于
log.Println("Error: something went wrong")
```

> 这条消息会写入错误输出流（stderr），即使程序的正常输出被重定向，错误信息也仍会显示在终端

---

## 程序信号控制（配合 `os/signal` 包）

用于优雅退出（如 Ctrl+C）：

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	fmt.Println("press Ctrl+C exit...")
	<-c // 阻塞等待信号
	fmt.Println("get exit signal, existing...")
}

```
