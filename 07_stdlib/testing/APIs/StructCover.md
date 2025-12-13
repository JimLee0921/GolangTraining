# `testing.Cover`

了解即可，代表一次测试运行的覆盖率全局快照，表示在某次 `go test -cover*` 运行中，哪些源文件被插桩、每个插桩点对应的计数器，以及使用了哪种覆盖率统计模式。

> 只有在使用 `go test -cover`（或其相关变体）时，才会产生覆盖率数据

## 定义

```
type Cover struct {
	Mode            string
	Counters        map[string][]uint32
	Blocks          map[string][]CoverBlock
	CoveredPackages string
}
```

- Mode：如何计算覆盖率，对应在使用 `go test -covermode=set|count|atomic`
    - set 为是否被执行过
    - count 为执行次数，非并发安全
    - atomic 为并发安全的执行测试，最准确
- Blocks：有哪些可统计的代码块，参考 [StructCoverBlock.md](StructCoverBlock.md)
- Counters：这些代码块实际被执行了多少次
