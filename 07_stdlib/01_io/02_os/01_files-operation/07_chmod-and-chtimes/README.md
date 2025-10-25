## 权限和时间戳修改

文件除了内容和名字，还有权限（读写执行）与时间属性（创建/访问/修改时间）

| 类型             | 概念            | 常用方法           |
|----------------|---------------|----------------|
| 权限（Permission） | 控制谁能读、写、执行    | `os.Chmod()`   |
| 时间戳（Timestamp） | 文件最后访问 / 修改时间 | `os.Chtimes()` |

### `os.Chmod`

用于修改文件权限。
语法：`os.Chmod(mode FileMode) error`

```
err := os.Chmod("script.sh", 0755)
if err != nil {
    panic(err)
}
```

这会将 script.sh 的权限改为：`rwxr-xr-x`

- 拥有者：读、写、执行
- 同组用户：读、执行
- 其他用户：读、执行

### `os.Chtimes`

文件系统中常见的两个时间：

| 英文        | 含义                  | 备注        |
|-----------|---------------------|-----------|
| **atime** | Access Time（最后访问时间） | 文件被读时更新   |
| **mtime** | Modify Time（最后修改时间） | 文件内容被改时更新 |

`os.Chtimes(name string, atime time.Time, mtime time.Time) error` 用于手动修改这两个时间

```
now := time.Now()
err := os.Chtimes("report.txt", now, now)
if err != nil {
    panic(err)
}
fmt.Println("修改时间戳成功")
```

> 在使用 os.Stat() 获取 os.FileInfo 打印的 info 是旧快照，一次运行需要重新读取才能看到效果