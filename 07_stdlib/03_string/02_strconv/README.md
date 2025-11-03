# strconv

strconv 是 Go 标准库中 用于字符串与数字/布尔类型互相转换的包。

可以把它理解为： 字符串（String） <-> 数字 / 布尔 / 字节 的转换工具集合

`str conv` 的意思就是：`String Convert`（字符串转换）


> 文档地址：https://pkg.go.dev/strconv


实际开发中：

* HTTP 请求参数是字符串
* JSON 字段中数字有时也是字符串
* 数据库读出来的数据经常需要转换
* 用户输入永远是字符串

而程序内部要求：

* id 是 int
* 价格是 float64
* 开关标志是 bool

所以必须掌握字符串与基础类型的转换

strconv 是 Web / API / JSON / 数据库 代码的高频必用包

| 类型          | Parse        | Format        | 快捷版         |
|-------------|--------------|---------------|-------------|
| bool        | ParseBool    | FormatBool    | -           |
| int (任意进制)  | ParseInt     | FormatInt     | -           |
| uint (任意进制) | ParseUint    | FormatUint    | -           |
| float       | ParseFloat   | FormatFloat   | -           |
| 十进制 int     | ParseInt(10) | FormatInt(10) | Atoi / Itoa |