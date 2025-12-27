# atomic.Value

用于在并发环境下，无锁地“整体发布/替换/读取”一个值（通常是一份不可变快照），用于读多写少的场景下，无锁发布与读取一份完整快照

```
type Value struct {
	// contains filtered or unexported fields
}
```

可以把它理解为：

* 读：`Load()` 无锁、非常快
* 写：`Store()` 原子替换
* 读到的是某次 `Store/Swap/CAS` 完成后的完整值

典型用途：

* 配置热更新（Config snapshot）
* 路由表、规则表整体替换
* 读多写少的全局状态快照

## 语义边界

`Value` 保证的是值本身的原子替换，不保证存进去的对象内部并发安全，所以主要适用于不可变快照：

* 写方：构造新对象（或深拷贝）-> `Store`
* 读方：`Load` 后只读使用
* 不要原地修改已发布对象

另外 `Store` 有强约束：类型必须一致

`atomic.Value` 有两条硬规则：

1. 第一次 `Store` 不能是 `nil`
2. 第一次 `Store` 决定了类型，之后所有 `Store` 必须是同一动态类型

例如：

```
var v atomic.Value
v.Store(&Config{}) // 第一次：*Config
v.Store(&Config{}) // 同类型
v.Store(Config{}) // 不同动态类型，会 panic
v.Store(nil)      // 也会 panic
```

> 这里的类型指的是 interface 里的动态类型，不是接口类型

## 核心方法

### 1. Load()

无锁读取当前值

```
func (v *Value) Load() (val any)
```

- 返回最近一次成功 `Store/Swap/CAS` 的值
- 如果从未 Store 过，返回 `nil`
- `Load` 返回 `any`，通常要类型断言

### 2. Store()

原子发布/替换值

```
func (v *Value) Store(val any)
```

- 原子地设置为 `val`
- 对后续 `Load()` 可见
- 必须满足上面的两条规则（非 nil + 类型一致）

### 3. Swap()

原子替换并返回旧值

```
func (v *Value) Swap(new any) (old any)
```

- 原子地把值替换为 `new`
- 返回旧值
- 同样受非 nil + 类型一致约束（相对于第一次 Store 的类型）

### 4. CompareAndSwap()

条件替换（CAS）

```
func (v *Value) CompareAndSwap(old, new any) (swapped bool)
```

- 只有当当前值与 `old` 相等时，才原子地替换为 `new`。
- 这里的相等是 interface 的比较规则，因此要求：
    - `old` 和当前值的动态类型必须是可比较的（否则会 panic）
    - 通常工程里用 CAS 时，传指针最稳（指针可比较）

## 对比 `atomic.Pointer[T]`

- 新代码、有泛型：优先使用 `atomic.Pointer[T]`
    - 类型安全
    - 支持 nil
    - CAS/Swap 更直观
- 需要存非指针值或兼容老代码：用 `atomic.Value`

`atomic.Value` 是无锁发布快照的通用盒子，`atomic.Pointer[T]` 是类型安全的指针版