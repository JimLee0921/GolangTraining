## 判断存在 / 获取信息

场景：启动时检查配置文件、确保目录已创建、按类型分支处理。

### `os.Stat`

`os.Stat(name string) (FileInfo, error)`

返回描述文件的 `os.FileInfo` 结构（大小、是否目录、权限、修改时间），如果发生错误，则返回 `*PathError` 类型

### `os.Lstat`

`os.Lstat(name string) (FileInfo, error)`
跟 Stat 类似（了解即可），但不跟随符号链接（symlink）

- 如果该文件是符号链接，则返回的 FileInfo 描述该符号链接
- Lstat 不会尝试跟踪该链接。如果发生错误，则返回 `*PathError` 类型的错误

> 在 Windows 上，如果文件是另一个命名实体（例如符号链接或已安装的文件夹）的代理重解析点，则返回的 FileInfo
> 会描述重解析点，并且不会尝试解析它

### `errors.Is(err, os.ErrNotExist) / os.IsNotExist(err error)bool`

判断不存在
