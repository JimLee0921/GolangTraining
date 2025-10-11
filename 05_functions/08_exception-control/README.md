## 异常控制机制

包括：`defer`、`panic`、`recover`

---

## 概览

在 Go 中，异常控制（exception control）由三部分组成：

| 机制        | 作用              | 类似概念（其他语言）        |
|-----------|-----------------|-------------------|
| `defer`   | 延迟调用（函数退出前执行）   | `finally`         |
| `panic`   | 主动触发运行时异常       | `throw` / `raise` |
| `recover` | 捕获 panic，恢复程序执行 | `try-catch`       |

它们共同构成 Go 的 **异常传播机制（stack unwinding）**

> Go 不鼓励滥用 panic，它只用于**不可恢复的错误**，一般业务逻辑应使用 `error` 类型返回值


> 官方BLOG文档：https://go.dev/blog/defer-panic-and-recover
>
> 官方WIKI文档：https://go.dev/wiki/PanicAndRecover

