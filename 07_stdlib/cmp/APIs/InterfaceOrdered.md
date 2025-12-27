# `cmp.Ordered`

接口 Ordered 是一个约束类型，只用于泛型，表示任何支持 `< <= >= >` 运算符的类型，
这种接口不是为了实现方法，而是为了限制类型。

```
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
```

`|` 为并集约束，意思是 `T` 必须是这些类型之一，`~` 是底层类型，表示除了这些 Go 原生类型，还可以是这些类型的自定义类型，并且
Go 希望可以使用自定义类型，比如：

```
int
type UserID int
type Price float64
```

这些都是满足 `~int` 的，如果没有 `~`，这些自定义类型就是不被允许的，只能匹配到 int