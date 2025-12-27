# cmp 包

cmp 包是 Go 标准库中用于比较（comparison）的工具包，主要目标是以一致、可组合、可拓展的方式比较两个值，并给出相等/小于/大于的结论。
并不是简单的 `==` 替代品，而是一个比较框架。

## cmd 之前比较

### 相等性比较

没有 cmp 包之前比较只能使用 `==` 进行比较：

```
if a == b {
    ...
}
```

但是 `==` 比较方式有较大弊端：

- 不支持 `slice` / `map` / `func`
- 不能自定义比较规则（比如忽略字段，允许误差）

### 排序比较

早期对于排序，比如 slice 排序，可以使用：

```
sort.Slice(users, func(i, j int) bool {
    return users[i].Age < users[j].Age
})
```

问题在于：

- 只有是否小于
- 没有统一的三态比较结果
- 逻辑分散不可复用

### 手写 compare 函数

```
func compare(a, b User) int {
    if a.Age < b.Age {
        return -1
    }
    if a.Age > b.Age {
        return 1
    }
    return 0
}
```

但是这样还有问题就是样板代码多，不够统一，不好组合多个字段

## cmp 包核心

cmp 包主要就是为了解决老版本哪些问题，把比较统一为三态结果（three-way comparison）

| 返回值  | 含义     |
|------|--------|
| `-1` | a < b  |
| `0`  | a == b |
| `+1` | a > b  |

> 这是计算机科学中非常经典的设计（C++、Java、Python 内部都大量使用）