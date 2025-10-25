# GO os 包

os 包主要提供操作操作系统提供的功能，比如：文件、目录、环境变量、进程、信号等。
虽然有些操作系统特定（Windows、Linux 差别），但 os 包对外接口是跨平台的。
如果需要更底层系统调用（syscall）或 OS 特定功能，则用 syscall 包或其它。

## 文件与目录相关

### `FileInfo` 对象

`os.FileInfo` 对象简化版

```
type FileInfo interface {
    Name() string       // 文件名（不含路径）
    Size() int64        // 文件大小（字节）如果是文件夹大小和平台相关（通常为 0 或 4096），不表示目录内容大小
    Mode() FileMode     // 文件权限和类型（可判断是否目录）
    ModTime() time.Time // 最后修改时间
    IsDir() bool        // 是否是目录
    Sys() any           // 底层系统数据（一般不用）
}
```

> 这其实是一个接口（interface），不是结构体。 `os.Stat()` 等方法返回的是真实类型（例如 *fs.fileStat）——但只关心它实现的这些方法

### `File` 对象

创建或读取文件时返回的结构体类型，定义在 os 包中：

```
type File struct {
    *file // 底层封装
}
```

最常用的不是结构体字段，而是它提供的方法。
**常用方法**

| 方法                                      | 功能              |
|-----------------------------------------|-----------------|
| `Write([]byte)` / `WriteString(string)` | 写入内容            |
| `Read([]byte)`                          | 读取内容            |
| `Seek(offset, whence)`                  | 改变读写指针位置        |
| `Close()`                               | 关闭文件            |
| `Name()`                                | 返回文件名           |
| `Stat()`                                | 获取文件信息          |
| `Sync()`                                | 强制将缓冲区数据刷入磁盘    |
| `Readdir(count int)`                    | 读取目录项（如果打开的是目录） |

### `DirEntry` 对象

DirEntry = fs.DirEntry，DirEntry 是从目录中读取的条目（使用ReadDir函数或File.ReadDir方法）

**常用方法**

| 方法        | 说明                     |
|-----------|------------------------|
| `Name()`  | 文件/目录名                 |
| `IsDir()` | 是否目录                   |
| `Type()`  | 返回 `fs.FileMode`       |
| `Info()`  | 获取完整 `FileInfo`（需额外调用） |

## Unix 文件权限

在 Go（以及 Linux/Unix）中，文件权限通常用八进制数字（以 0 开头）表示。

0755 是八进制权限表示（octal file mode） 表示

0755 是一个常见的权限值，常用于文件夹或可执行文件。

- 拥有者：读写执行 (7)
- 组用户：读执行 (5)
- 其他用户：读执行 (5)

### 分解结构

| 位数  | 代表对象          | 值   | 含义                            |
|-----|---------------|-----|-------------------------------|
| 第1位 | 特殊位（一般省略）     | `0` | 通常不用，保留给 SUID/SGID/Sticky bit |
| 第2位 | 所有者 (user)    | `7` | `rwx`（读写执行）                   |
| 第3位 | 组 (group)     | `5` | `r-x`（读执行）                    |
| 第4位 | 其他用户 (others) | `5` | `r-x`（读执行）                    |

### 二进制含义

| 权限字符          | 二进制   | 八进制值 | 说明  |
|---------------|-------|------|-----|
| `r` (read)    | `100` | 4    | 可读  |
| `w` (write)   | `010` | 2    | 可写  |
| `x` (execute) | `001` | 1    | 可执行 |

### go 使用

```
os.Mkdir("dir", 0755)   // 创建目录设置权限
os.Chmod("file", 0755)  // 修改权限
```

> Windows 会忽略这些 Unix 权限值，但仍需传入

```temp
权限与时间（os.Chmod / os.Chtimes）
```

## 常量补充

`os.PathSeparator` 和 `os.PathListSeparator` 是系统路径语法相关的底层常量

```
const (
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator
)
```

| 常量名                      | 类型   | Linux/macOS 值 | Windows 值 | 说明          |
|--------------------------|------|---------------|-----------|-------------|
| `filepath.Separator`     | rune | `'/'`         | `'\\'`    | 文件夹内路径分隔符   |
| `filepath.ListSeparator` | rune | `':'`         | `';'`     | 环境变量路径列表分隔符 |

> Separator 用来分隔路径内部的层级，而 ListSeparator 用来分隔多个路径之间。
