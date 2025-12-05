## 7️⃣ 集合：Struct 字段 / 方法

### 字段访问

* **`(v Value) NumField() int`**
  struct 字段数。
* **`(v Value) Field(i int) Value`**
  按 index 取字段，类似 `v.Field(i)`。
* **`(v Value) FieldByName(name string) Value`**
  按字段名查；找不到返回 invalid Value。
* **`(v Value) FieldByIndex(index []int) Value`**
  支持嵌套字段：`[]int{a,b,c}` 相当于 `Field(a).Field(b).Field(c)`。
* **`(v Value) FieldByIndexErr(index []int) (Value, error)`**
  和上面类似，但返回 error，便于区分错误。
* **`(v Value) FieldByNameFunc(match func(string) bool) Value`**
  用回调匹配字段名（比如不区分大小写匹配）。

### 方法访问

* **`(v Value) NumMethod() int`**
  导出方法数量（和 Type.NumMethod 规则类似）。
* **`(v Value) Method(i int) Value`**
  按 index 取方法，返回的是一个“可调用的 Value”（Kind 为 Func）。
* **`(v Value) MethodByName(name string) Value`**
  按方法名查找；找不到返回 invalid Value。

> 学习要点：`Field*` 和 `Method*` 都只对 **Kind=Struct 或带方法的类型** 有意义。

---

## 9️⃣ 函数 / 方法调用

### 调用现有函数 / 方法

* **`(v Value) Call(in []Value) []Value`**
  反射调用函数或方法。`in` 是参数 `Value` 列表。返回也是 `Value` 列表。
* **`(v Value) CallSlice(in []Value) []Value`**
  专门用于最后一个参数是变长参数（...T）的情况。

> 这两个的前提：`v.Kind() == Func` 或 `v` 是 Method 得到的函数值。

### 生成“动态函数”

* **`MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value`**
  构造一个“函数 Value”，它的类型是 `typ`，底层调用你提供的 `fn`。
  常见用途：**动态代理、打日志、mock、hook** 等。



