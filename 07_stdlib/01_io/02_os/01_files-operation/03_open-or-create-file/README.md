## 文件创建与打开

| 函数                              | 功能          | 是否自动创建 | 可写入      | 是否清空原内容  |
|---------------------------------|-------------|--------|----------|----------|
| `os.Create(name)`               | 创建新文件       | 是      | 是        | 是        |
| `os.Open(name)`                 | 打开已存在文件（只读） | 否      | 否        | 否        |
| `os.OpenFile(name, flag, perm)` | 自定义方式打开     | 可选     | 取决于 flag | 取决于 flag |

文件操作的核心动作，几乎所有写日志、存储配置、导出数据的功能都要用到

### `os.Create`

`os.Create(name) (*File, error)`

- 创建文件（若存在则清空）
- 文件不存在->自动创建；文件已存在->清空内容再写
- 默认权限：0666（受 umask 影响）底层为 `os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)`

### `os.Open`

`os.Open(name string) (*File, error)`
打开只读文件，用于读取内容

### `os.OpenFile`

自定义打开方式（最常用），更加灵活，可以进行追加，读写等操作

`os.OpenFile(name string, flag int, perm FileMode) (*File, error)`

参数：

- name：路径
- flag：打开方式（位标志，可用 | 组合），控制行为
- perm：权限（仅在创建文件时生效；Unix 有效，Windows基本忽略），控制文件权限

返回值是 *os.File（实现 io.Reader、io.Writer、io.Seeker、io.Closer 等），用它来 `Read/Write/Seek/Close/Sync`

**flag**
底层是 syscall 定义的

```
const (
    O_RDONLY int = syscall.O_RDONLY // 只读
    O_WRONLY int = syscall.O_WRONLY // 只写
    O_RDWR   int = syscall.O_RDWR   // 读写
    O_APPEND int = syscall.O_APPEND // 追加写入
    O_CREATE int = syscall.O_CREAT  // 不存在则创建
    O_EXCL   int = syscall.O_EXCL   // 与 O_CREATE 一起用，若已存在则报错
    O_SYNC   int = syscall.O_SYNC   // 同步写入磁盘
    O_TRUNC  int = syscall.O_TRUNC  // 打开时清空文件内容
)
```

| 标志            | 含义                                  |
|---------------|-------------------------------------|
| `os.O_RDONLY` | 只读打开                                |
| `os.O_WRONLY` | 只写打开                                |
| `os.O_RDWR`   | 读写打开                                |
| `os.O_CREATE` | 文件不存在则创建                            |
| `os.O_TRUNC`  | 打开后将文件截断为 0（清空）                     |
| `os.O_APPEND` | 追加写（每次写都追加到末尾）                      |
| `os.O_EXCL`   | 与 `O_CREATE` 一起用：若文件已存在则**报错**，防止覆盖 |

`flag` 使用按位或运算符`|`进行组合使用 `f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)`

### `os.CreateTemp`

`os.CreateTemp(dir, pattern string) (*File, error)`

- `dir 参数`：指定在哪个目录下创建临时目录，如果传入空字符串 ""，Go 会自动使用系统默认的临时目录（例如 /tmp 或
  %TEMP%）
- `pattern 参数`：临时目录名称的前缀。Go 会在它后面自动加上一串随机字符串，确保唯一性

在指定目录中创建临时文件（自动随机文件名），常用于下载中间文件、缓存文件、测试用