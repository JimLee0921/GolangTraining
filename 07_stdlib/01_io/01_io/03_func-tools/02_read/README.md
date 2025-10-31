## `io.ReadAll` / `io.ReadFull`

| 函数                        | 停止条件             | 场景               | 注意点              |
|---------------------------|------------------|------------------|------------------|
| **`io.ReadAll(r)`**       | **直到 EOF**       | 读取“整个流”          | **可能一次把大文件读爆内存** |
| **`io.ReadFull(r, buf)`** | **直到 buf 满** 或报错 | 想读取**完全固定长度的数据** | 若不足会返回错误         |

### `io.ReadAll`

```
func ReadAll(r Reader) ([]byte, error)
```

直接把 r 的全部数据一次性读进内存。

适合：

- 读小文件
- 读 HTTP Response Body
- 读内存流

> ReadAll 会把数据读进内存，如果数据很大可能直接 OOM（内存打爆）
>
> 小文件 -> 可以用 ReadAll
>
> 大文件 / 网络流 -> 用 io.Copy 或 `bufio.Reader`


### `io.ReadFull`

```
func ReadFull(r Reader, buf []byte) (n int, err error)
```

尝试填满整个 buf，直到读满或遇到错误。

适合：

- 二进制协议
- 网络分包读取
- 从文件中读取固定长度结构体