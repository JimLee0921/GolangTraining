## 快捷读写文件

`os.ReadFile` / `os.WriteFile`
这两个函数是对文件读写的高层封装，是 Go 官方推荐的简化方案，特别适合处理配置、日志、小型数据文件等。

如果是使用 file 的那些方法

```
f, _ := os.Open()
f.Read(buf)
f.Close()
```

要自己管理文件句柄、缓冲区、关闭资源，而

```
data, _ := os.ReadFile("file.txt")
```

一步写入，简单、可靠、最常用。仅适合小文件（内存可一次装下的）

### `os.ReadFile`

`ReadFile(name string) ([]byte, error)`，用于一次性读取整个文件，
返回内容（[]byte）和错误（error）
特点:

- 文件不存在 -> 报错
- 返回 []byte 类型，可直接转成字符串
- 自动打开 -> 读取 -> 关闭（内部已封装）

### `os.WriteFile`

`os.WriteFile(name string, data []byte, perm FileMode) error`，用于一次性写入文件，每次调用都会重写整个文件（覆盖写）

**参数：**

| 参数     | 类型            | 说明             |
|--------|---------------|----------------|
| `name` | `string`      | 文件名（路径）        |
| `data` | `[]byte`      | 要写入的内容         |
| `perm` | `fs.FileMode` | 文件权限（如 `0644`） |

**特点：**

- 文件不存在 -> 自动创建
- 文件存在 -> 会被覆盖
- 需要权限值：如 0644
- 自动打开 -> 写入 -> 关闭

> 可以配合 ReadFile + WriteFile 来实现追加等功能，更推荐使用 os.OpenFile + O_APPEND
