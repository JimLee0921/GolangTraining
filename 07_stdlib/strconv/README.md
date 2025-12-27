# strconv

strconv 包主要用于基本数据类型与字符串表示形式之间的转换。可以理解为：字符串（String） <-> 数字 / 布尔 / 字节


> 文档地址：https://pkg.go.dev/strconv


在 Go 中字符串不是数据，只是表现形式，而程序真正要计算、比较、存储的是数值/布尔值/二进制/Unicode code point

比如：

| 场景                      | 输入/输出          |
|-------------------------|----------------|
| HTTP / JSON / CLI / Env | string         |
| 配置文件 / CSV / 日志         | string         |
| 数据库存储                   | string / bytes |
| 用户输入                    | string         |
| 程序计算                    | 强类型            |

而 strconv 就是为了处理这些数据，面向程序进行字符串处理