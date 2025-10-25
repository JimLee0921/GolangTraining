## 创建目录

> perm 参数见下面 unix 文件权限

### `os.Mkdir`

- `os.Mkdir(name string，perm FileMode) error`
- 创建单层目录，如果目录已存在或父目录不存在会报错

### `os.MkdirAll`

- `os.MkdirAll(name string，perm FileMode) error`
- 创建多级目录，递归创建父级
- 自动创建所有不存在的父目录，如果路径已存在，不会报错，是生产环境中最常用的版本

### `os.MkdirTemp()`

`os.MkdirTemp(dir, pattern string) (name string, err error)`

- `dir 参数`：指定在哪个目录下创建临时目录，如果传入空字符串 ""，Go 会自动使用系统默认的临时目录（例如 /tmp 或
  %TEMP%）
- `pattern 参数`：临时目录名称的前缀。Go 会在它后面自动加上一串随机字符串，确保唯一性

> 创建临时目录，通常在需要生成临时文件夹（例如缓存、测试、临时下载等）时使用

| 方法               | 作用     | 特点       | 返回值              |
|------------------|--------|----------|------------------|
| `os.Mkdir()`     | 创建单层目录 | 父目录必须已存在 | error 错误信息       |
| `os.MkdirAll()`  | 创建多级目录 | 安全且常用    | error 错误信息       |
| `os.MkdirTemp()` | 创建临时目录 | 自动随机名    | 临时文件目录+error错误信息 |
