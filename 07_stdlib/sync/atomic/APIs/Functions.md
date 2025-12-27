# atomic 顶层函数

`sync/atomic` 顶层函数本质上是按原子操作原语来分的；不同类型（Int32/Int64/Uint32/Uint64/Uintptr/Pointer）只是同一类操作的不同变体，

这些顶层函数函数是旧接口，`atomic.Int32 / Int64 / Uint64 / Pointer[T]` 是新接口，新版本项目推荐使用新接口，这里不再讲解，使用方法在各自新接口都是对应的

## 1. 原子加法类（Fetch-Add，返回新值）

用于并发计数器/序号/累加统计

- `AddInt32(addr *int32, delta int32) (new int32)`
- `AddInt64(addr *int64, delta int64) (new int64)`
- `AddUint32(addr *uint32, delta uint32) (new uint32)`
- `AddUint64(addr *uint64, delta uint64) (new uint64)`
- `AddUintptr(addr *uintptr, delta uintptr) (new uintptr)`

## 2. 原子按位操作类（Bitwise，返回旧值）

用于并发 flags（位标志）：置位用 Or，清位/筛选用 And

- AND（返回 old）

    - `AndInt32(addr *int32, mask int32) (old int32)`
    - `AndInt64(addr *int64, mask int64) (old int64)`
    - `AndUint32(addr *uint32, mask uint32) (old uint32)`
    - `AndUint64(addr *uint64, mask uint64) (old uint64)`
    - `AndUintptr(addr *uintptr, mask uintptr) (old uintptr)`

- OR（返回 old）

    - `OrInt32(addr *int32, mask int32) (old int32)`
    - `OrInt64(addr *int64, mask int64) (old int64)`
    - `OrUint32(addr *uint32, mask uint32) (old uint32)`
    - `OrUint64(addr *uint64, mask uint64) (old uint64)`
    - `OrUintptr(addr *uintptr, mask uintptr) (old uintptr)`

## 3. 原子比较并交换类（CAS：Compare-And-Swap，返回是否成功）

用于无锁状态机、乐观并发更新、只允许一次成功等

- `CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)`
- `CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)`
- `CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)`
- `CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)`
- `CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)`
- `CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)`

## 4. 原子读取类（Load）

无锁读取一个共享变量的当前值（保证可见性）。

- `LoadInt32(addr *int32) (val int32)`
- `LoadInt64(addr *int64) (val int64)`
- `LoadUint32(addr *uint32) (val uint32)`
- `LoadUint64(addr *uint64) (val uint64)`
- `LoadUintptr(addr *uintptr) (val uintptr)`
- `LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)`

## 5. 原子写入类（Store）

无锁发布一个共享变量的新值（保证后续 Load 可见）。

- `StoreInt32(addr *int32, val int32)`
- `StoreInt64(addr *int64, val int64)`
- `StoreUint32(addr *uint32, val uint32)`
- `StoreUint64(addr *uint64, val uint64)`
- `StoreUintptr(addr *uintptr, val uintptr)`
- `StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)`

## 6. 原子交换类（Swap，返回旧值）

用于“设置新值并拿到旧值”，常见模式：**取出并清零、门闩、切换指针**。

- `SwapInt32(addr *int32, new int32) (old int32)`
- `SwapInt64(addr *int64, new int64) (old int64)`
- `SwapUint32(addr *uint32, new uint32) (old uint32)`
- `SwapUint64(addr *uint64, new uint64) (old uint64)`
- `SwapUintptr(addr *uintptr, new uintptr) (old uintptr)`
- `SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)`

## 7. 指针专用子集（unsafe.Pointer 相关）

它们属于上面的 CAS/Load/Store/Swap 四类，只是操作对象是 `*unsafe.Pointer`：

- `LoadPointer`
- `StorePointer`
- `SwapPointer`
- `CompareAndSwapPointer`


