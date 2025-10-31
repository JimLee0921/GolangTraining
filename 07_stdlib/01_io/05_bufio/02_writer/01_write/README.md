## Write 写入

`bufio.Writer` 的最底层、最核心方法，所有其它写入函数（包括 `WriteString`、`WriteRune` 等）本质上都基于它。

### 方法签名

```
func (b *Writer) Write(p []byte) (nn int, err error)
```

参数：

- p: 字节切片，需要写入的数据，Write() 不会修改 p 的内容，会把数据尽可能多地写入缓冲区中

返回值：

- nn：实际写入缓冲区的字节数
    - 表示有多少字节成功写入缓冲区
    - 如果写入一部分就出错，则 nn < len(p)
    - 如果完全写入，则 nn == len(p)
- err：错误（如果底层写入失败）
    - 如果写入完全成功 -> `err == nil`
    - 如果底层写入失败（例如磁盘满、连接断开） -> `err != nil`
    - 即使出错，也可能部分写入（nn > 0）

### 注意事项

1. 虽然缓冲区满时会自动调用 `Flush`，但是仍然需要手动 `Flush`，否则最后的内容无法真正写入
2. 如果是写入文件，Flush 必须发生在 Close 之前。也可以使用 defer 的 LIFO（后登记先执行）规则：先 defer f.Close()，再 defer
   w.Flush()，这样退出时会先 Flush 再 Close