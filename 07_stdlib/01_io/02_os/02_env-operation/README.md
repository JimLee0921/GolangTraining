# os 包环境变量相关

环境变量是操作系统级的键值对，例如：

```
PATH=/usr/bin:/usr/local/bin
HOME=/home/user
GOPATH=/Users/lee/go
```

* `PATH`：系统可执行程序路径
* `HOME`：当前用户主目录
* `GOPATH`：Go 执行目录
* `TEMP`：临时目录
* 自定义变量（如 `APP_ENV=prod`）

Go 程序可以通过 os 包来读取、设置、删除这些变量。

## env 相关

### 函数总览

| 函数                      | 功能                              | 说明  |
|-------------------------|---------------------------------|-----|
| `os.Getenv(key)`        | 获取环境变量值（不存在返回空字符串）              | 常用  |
| `os.LookupEnv(key)`     | 获取环境变量值（区分变量是否存在）               | 更安全 |
| `os.Setenv(key, value)` | 设置（或修改）环境变量                     |     |
| `os.Unsetenv(key)`      | 删除环境变量                          |     |
| `os.Clearenv()`         | 清空所有环境变量                        | 慎用  |
| `os.Environ()`          | 返回所有环境变量（形式为 `"KEY=value"` 的切片） |     |
| `os.Expand(s, mapping)` | 根据自定义规则展开变量（${VAR}）             |     |
| `os.ExpandEnv(s)`       | 展开字符串中的环境变量                     |     |

### 1. 获取环境变量：`os.Getenv`

语法：`Getenv(key string) string`

```
path := os.Getenv("PATH")
fmt.Println("系统 PATH:", path)
```

**特点：**

* 用于检索由键指定的环境变量的值。它返回该值，如果该变量不存在，则返回空值
* 要区分空值和未设置的值，可以使用LookupEnv。
* 不区分大小写（Windows），区分大小写（Linux/macOS）

---

### 2. 判断环境变量是否存在：`os.LookupEnv`

语法：`LookupEnv(key string ) ( string , bool )`

```
value, exists := os.LookupEnv("HOME")
if exists {
fmt.Println("HOME:", value)
} else {
fmt.Println("HOME 未定义")
}
```

* 检索由键指定的环境变量的值
* 如果该变量存在于环境中，则返回值（可能为空）且布尔值为 true，否则返回值为空且布尔值为 false
* 比 `os.Getenv` 更安全

---

### 3. 设置环境变量：`os.Setenv`

语法：`Setenv(key, value string) error`

```
os.Setenv("APP_MODE", "debug")
fmt.Println("APP_MODE =", os.Getenv("APP_MODE"))
```

* 修改只对当前进程及其子进程有效
* 程序退出后环境变量不会保留到系统中

---

### 4. 删除指定环境变量：`os.Unsetenv`

语法：`Unsetenv(key string) error`

```
os.Unsetenv("APP_MODE")
```

* 删除指定的环境变量
* 等价于 Linux 的 unset KEY
* 若变量不存在，不会报错
* 仅影响当前进程及其子进程。

---

### 5. 清空所有环境变量 `os.Clearenv()`

小心使用：这会让系统变量（如 PATH）消失，影响后续操作。

- 当前进程的环境变量表会被清空
- 调用 `os.Getenv()` 或 `os.Environ()` 都会返回空结果
- 所有之后启动的 子进程（exec.Command） 都不会继承任何环境变量
- 无法恢复，除非手动重新设置

### 6. 列出所有环境变量：`os.Environ`

Environ 以`"key=value"`的形式返回代表环境的字符串的副本。

语法：`os.Environ() []string`

```
for _, e := range os.Environ() {
    fmt.Println(e)
}
```

输出格式：

```
PATH=/usr/bin:/bin
HOME=/home/user
LANG=en_US.UTF-8
```

---

### 7. 自定义变量替换规则 `os.Expand`

自定义变量展开

语法：`Expand(s string, mapping func(string) string) string`

| 参数        | 类型                    | 说明                                |
|-----------|-----------------------|-----------------------------------|
| `s`       | `string`              | 包含变量占位符的字符串，例如 `"Hello, ${NAME}"` |
| `mapping` | `func(string) string` | 自定义函数：给定变量名，返回它的值                 |

- `os.Expand` 会扫描字符串 s，识别 `$var` 和 `${var}` 模式
- 对每个变量名，会调用 `mapping(key)` 并用返回值替换变量
- 若 mapping 返回空字符串，则变量会被清空

| 场景                 | 举例                       |
|--------------------|--------------------------|
| 自定义模板替换            | `${PROJECT}` -> `Shopee` |
| 程序内部配置变量           | `$DB_HOST` -> 配置表中的主机地址  |
| 读取 `.env` 或自定义 map | 从 map 中动态替换              |

### 8. 自动使用系统环境变量 `os.ExpandEnv()`

根据当前环境变量的值替换字符串中的 `${var}` 或 `$var`。对未定义变量的引用将被替换为空字符串。

语法：`ExpandEnv(s string) string`

等同于：`os.Expand (s, os.Getenv)`，内部自动用当前系统环境变量作为 mapping。

| 特点                        | 说明             |
|---------------------------|----------------|
| 自动读取系统环境变量                | 不需要自己写 mapping |
| 未定义变量                     | 替换为空字符串        |
| 支持 `$VAR` 与 `${VAR}` 两种形式 | 与 Shell 一致     |

