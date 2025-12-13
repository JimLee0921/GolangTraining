# `testing.PB`

`testing.PB` 是 `Parallel Benchmark` 的缩写，用于 `B.RunParallel` 内部，作为每个 worker goroutine 的迭代器，用于控制还要不要继续下一次操作

- `testing.B` 负责整个 benchmark 的控制
- 而 `testing.PB` 就是负责单个 worker goroutine 在并行场景下的本次循环控制

> 可以理解为并行版本的 `b.Loop` / `b.N` 分发器

## 定义

```
type PB struct {
	// contains filtered or unexported fields
}
```

没啥说的

## 方法

只有一个方法 Next，询问测试框架：这一轮我还需不需要再执行一次操作

- 如果需要，就给一个配额，返回 true
- 如果已经没配额了，就返回 false

```
func (pb *PB) Next() bool
```

## 使用讲解

典型用法如下：

```
b.RunParallel(func(pb *testing.PB) {
	for pb.Next() {
		// 并发执行的 benchmark 操作
		work()
	}
})
```

1. 每个并发 goroutine 拿到一个 `*PB`
2. 通过不断调用 `pb.Next()` 来索要下一次任务
3. `Next()` 返回 true 就执行一次 `work()` 操作
4. 返回 false 表示该 goroutine 可以退出了