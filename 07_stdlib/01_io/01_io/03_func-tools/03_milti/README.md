## `io.MultiReader` / `io.MultiWriter`

用来构建多个输入合成一个输出 / 一个输入分发到多个输出。

| 工具                              | 作用                         | 类比              |
|---------------------------------|----------------------------|-----------------|
| **`io.MultiReader(r1, r2, …)`** | 把多个 Reader 串成一个 Reader     | **把多个河道接成一条河流** |
| **`io.MultiWriter(w1, w2, …)`** | 把一个 Writer 的输出复制给多个 Writer | **把一条水流分成多条支流** |

### `io.MultiReader`

```
func MultiReader(readers ...Reader) Reader
```

按顺序依次读取多个 Reader，前一个读完（返回 EOF）后，自动切换到下一个

特点：

- 不会并发读取
- 是按顺序读取
- 遇到 EOF 会切到下一个

### `io.MultiWriter`

```
func MultiWriter(writers ...Writer) Writer
```

一次写入数据，自动写入所有 Writer。
如果其中一个 Writer 写失败：写错误会立刻返回，但部分 Writer 可能已经写成功，出现数据不一致风险

| 使用 MultiWriter 时 | 适用场景        |
|------------------|-------------|
| 写日志时             | 即使一个目标失败无所谓 |
| 广播调试信息           | 只是辅助用途      |
| 写数据库 / 写关键数据     | 要求原子性不建议使用  |
