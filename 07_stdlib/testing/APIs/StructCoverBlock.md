# `testing.CoverBlock`

**了解即可**，CoverBlock 描述的是源码中一段可以被单独统计是否执行过的代码块的位置与权重。
不是行，不是语句，是编译器切出来的基本块（basic block）

## 定义

```
type CoverBlock struct {
    Line0 uint32 // 起始行
    Col0  uint16 // 起始列
    Line1 uint32 // 结束行
    Col1  uint16 // 结束列
    Stmts uint16 // 该块包含的语句数
}
```

- `Line0 / Col0`：起点。这个代码块从哪一行、哪一列开始，行号、列号都是 `1-based`
- `Line1 / Col1`：终点。这个代码块在哪一行、哪一列结束，是源码位置范围，不是 AST 节点
- `Stmts`：权重，表示这个代码块里包含多少条可执行语句（不是行数，不是 AST 节点数，是 statement count）