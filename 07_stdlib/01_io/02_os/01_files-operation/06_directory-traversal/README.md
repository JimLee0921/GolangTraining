## 遍历目录

go 遍历指定目录有主要下面三种方式

| 方法                   | 来源          | 特点                            |
|----------------------|-------------|-------------------------------|
| `os.ReadDir()`       | Go 1.16+ 推荐 | 读取**单层目录**，简洁高效               |
| `filepath.Walk()`    | 旧接口         | 递归遍历目录（返回 `os.FileInfo`）      |
| `filepath.WalkDir()` | Go 1.16+ 推荐 | 递归遍历目录（返回 `fs.DirEntry`），性能更优 |

### `os.ReadDir`

只列出当前目录下的内容（不递归），返回按文件名排序的所有目录条目。如果读取目录时发生错误，ReadDir 将返回错误发生前能够读取的条目以及错误本身

语法：`ReadDir(name string) ([]DirEntry, error)`

### `filepaht.Walk()`

> 了解即可

语法：`func Walk(root string, walkFn filepath.WalkFunc) error`

WalkFunc 定义如下：
`type WalkFunc func(path string, info os.FileInfo, err error) error`

- path：当前遍历到的路径（文件或目录）
- info：os.FileInfo 对象（包含名称、大小、权限等）
- err：如果访问该文件出错，会传入错误
- 返回值：如果返回非 nil，遍历会中止

| 优点                 | 缺点                              |
|--------------------|---------------------------------|
| 易用，老版本兼容           | 每次都要调用 `os.Stat()` 获取文件信息（性能稍慢） |
| 回调传入 `os.FileInfo` | 无法直接跳过目录（要手动控制）                 |

### `filepath.WalkDir`

Go 1.16 新增 WalkDir，是 Walk 的优化版本。

语法：`func WalkDir(root string, fn fs.WalkDirFunc) error`

WalkDirFunc 定义如下：
`type WalkDirFunc func(path string, d fs.DirEntry, err error) error`

- d：类型为 `fs.DirEntry`，直接从 ReadDir 读取，无需再次调用 `os.Stat()`，`d.Info()` 返回 os.FileInfo