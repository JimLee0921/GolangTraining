## 删除/移动文件或文件夹

### `os.Remove`

`os.Remove(name string) error`，用于删除文件或空目录。

**特点：**

- 如果目标是普通文件 -> 直接删除
- 如果目标是空目录 -> 也能删除
- 如果是非空目录 -> 会报错：remove xxx: directory not empty

### `os.RemoveAll`

`os.RemoveAll(name string) error`，用于递归删除整个目录树（包含所有子目录和文件）。

**特点：**

- 可同时删除目录及其所有内容
- 已存在即删，不存在也不报错
- 非常危险 —— 类似 Linux 的 rm -rf。

常用于：清空缓存文件夹；单元测试清理；重置项目输出目录。

**删除注意事项**

| 场景      | 建议                                  |
|---------|-------------------------------------|
| 删除单个文件  | 用 `os.Remove()`                     |
| 删除空目录   | 仍用 `os.Remove()`                    |
| 删除整个目录树 | 用 `os.RemoveAll()`                  |
| 删除前确认路径 | 打印或检查（防误删）                          |
| 并发删除    | 小心 race condition（多个 goroutine 同时删） |

### `os.Rename`

可以重命名或者移动文件或文件夹，可以把它理解成 Linux 的 mv 命令，或者 Windows 的重命名/移动文件操作。
`os.Rename(oldname, newname string) error`

| 参数        | 含义                 |
|-----------|--------------------|
| `oldpath` | 原文件（或目录）的路径        |
| `newpath` | 新文件（或目录）的路径        |
| 返回值       | `error`（操作失败时返回错误） |

- 本质上是同一分区内移动文件
- os.Rename 在大多数操作系统上只能在同一个文件系统（分区）内工作，如果跨分区移动可能会报错

| 场景      | 说明                                      |
 |---------|-----------------------------------------|
| 目标文件已存在 | 会覆盖旧文件（部分系统可能拒绝）                        |
| 源文件不存在  | 报错：`no such file or directory`          |
| 跨分区移动   | 报错：`invalid cross-device link`（需用复制+删除） |

### `os.Truncate`

语法：`os.Truncate(name string, size int64) error`

将文件截断为指定大小（不是删除）`os.Truncate("log.txt", 0)` 清空内容但保留文件，不常用