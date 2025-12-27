# sort 包

sort 包提供了对切片和用户自定义集合进行原地排序的能力，并允许控制比较规则

> 切片原地排序且可自定义比较逻辑

## 常见用途

### 1. 对内置类型切片排序

```
sort.Ints([]int)
sort.Strings([]string)
sort.Float64s([]float64)
```

### 2. 对任意类型进行排序

通过实现一个 Interface 接口就可以排序任何数据结构：

```
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

- `[]struct`
- 自定义 slice

### 3. 自定义排序规则

可以按照多字段 / 倒序 / 业务规则进行排序

### 4. 可以判断是否已排序/稳定排序

主要用于数据管道/聚合/二次排序

```
sort.IsSorted(data)
sort.Stable(data)
```

## 设计思想

一切围绕一个接口：`sort.Interface`，在实现 Interface 接口后排序算法不关心数据是什么，只关心 `i / j` 谁应该在前。

并且排序算法事黑盒的，也就是不需要关心使用的快排还是堆排，是否递归，如何交换，只需要定义什么是小于即可。