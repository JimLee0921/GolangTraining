# cmp 顶层函数

主要就三个函数，这三个函数也是基于 type Ordered 约束的

## 1. Compare

专门为泛型比较和排序设计的工具函数，很多算法比如排序/二分查找/有序结构/多字段排序都需要 compare 函数而不是 bool，Compare 是
cmp 包的基石函数

```
func Compare[T Ordered](x, y T) int
```

**参数**

- `x,y T`：`T Ordered`，表示必须支持 `< <= >= >`

**返回值**

int 表示三态比较，注意不能严格保证 `-1 / 0 / 1`，只能依赖符号，而不能依赖具体的数值

| 返回值   | 含义     |
|-------|--------|
| `< 0` | x < y  |
| `0`   | x == y |
| `> 0` | x > y  |

**特殊值比较规则**

这是 `cmp.Compare` 最重要且与普通运算符不同的地方，标准浮点算法中，`NaN`(Not a Number) 与任何值比较（包括它本身）都会返回
false，这会导致排序算法等崩溃。

`cmp.Compare` 专门解决了这个问题：

- NaN 最小化原则：NaN 被认为比任何非 NaN 值都要小
- NaN 相等原则：`NaN == NaN` 在此函数中返回 0

## 2. Less

`cmp.Less` 是 `cmp.Compare` 的简化版（`cmp.Less(x, y) == (cmp.Compare(x, y) < 0)`），作用非常纯粹：判断第一个值是否严格小于第二个值。

如果说 Compare 返回的是整数`（-1, 0, 1）`，那么 Less 返回的就是布尔值（true 或 false）。

```
Less[T Ordered](x, y T) bool
```

> 表示在 `cmp.Compare` 定义的顺序下，x 是否小于 y

## 3. Or

返回参数中第一个非零的值，如果没有参数不为零，则返回零值

```
func Or[T comparable](vals ...T) T
```

- vals 的泛型约束 `T` 不是 Ordered，而是 comparable，是可变参数
- `T` 只需要 comparable 是因为判断零值用的是 `v != zero(T)`，这里需要 `==` 而不需要 `< <= >= >`