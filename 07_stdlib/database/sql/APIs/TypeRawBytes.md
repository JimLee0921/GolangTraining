# sql.RawBytes

`sql.RawBytes` 是一个零拷贝的 Scan 目标，用来直接引用数据库驱动内部缓冲区的数据，只用于数据库返回二进制/字节型数据时才有意义

> 太过底层，基本用不到，了解即可

```
type RawBytes []byte
```

## 对比普通 Scan

### 普通 Scan

安全但是有拷贝：

```
var b []byte

rows.Scan(&b)
```

- 驱动内部 buffer 拷贝一份 b
- b 的内存归调用方所有，可以长期保存、修改
- 安全但是有一次内存拷贝

### RawBytes Scan

不进行拷贝：

```
var rb sql.RawBytes

rows.Scan(&rb)
```

- 驱动内部 buffer 不进行拷贝
- rb 直接指向驱动内部内存
- 性能上更快，但是及其危险