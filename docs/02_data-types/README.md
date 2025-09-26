# Go 数据类型笔记

> Go 语言中常见数据类型，索引和导航如下

## 结构导航

- [原始类型](01_primitive-types/README.md)
    - [布尔类型](01_primitive-types/01_boolean.md)
    - [数值类型](01_primitive-types/02_numbers.md)
    - [字符串](01_primitive-types/03_string.md)
- [复合类型](02_composite-types/README.md)
    - [数组](02_composite-types/01_array.md)
    - [切片](02_composite-types/02_slice.md)
    - [映射](02_composite-types/03_map.md)
    - [结构体](02_composite-types/04_struct.md)
    - [指针](02_composite-types/05_pointer.md)
    - [函数](02_composite-types/06_function.md)
    - [通道](02_composite-types/07_channel.md)
    - [接口](02_composite-types/08_interface.md)`

> 也可以按照值类型 vs 引用类型进行划分

---

## 值类型 (Value Types)

赋值/传参时会完整复制一份新数据，变量之间相互独立。

* **原始类型 (Primitive types)**

    * 布尔：`bool`
    * 数值：`int`, `float64`, ...
    * 字符串：`string`（语义上表现为值类型，虽然底层实现有点特殊）

* **复合类型 (Composite types)**

    * 数组 `array`
    * 结构体 `struct`

---

## 引用类型 (Reference Types)

赋值/传参时复制的是“引用”，多个变量共享同一份底层数据，修改一处会影响另一处。

* **复合类型 (Composite types)**

    * 切片 `slice`
    * 映射 `map`
    * 指针 `pointer`
    * 函数 `function`（函数值本身是引用语义，多个变量可指向同一个函数体）
    * 通道 `channel`
    * 接口 `interface`（保存“动态类型 + 动态值”的引用信息）

---

## 对照总结表

| 分类维度     | 值类型                                | 引用类型                                                          |
|----------|------------------------------------|---------------------------------------------------------------|
| **原始类型** | `bool`, `int`, `float64`, `string` | （无）                                                           |
| **复合类型** | `array`, `struct`                  | `slice`, `map`, `pointer`, `function`, `channel`, `interface` |

---

总结：

* **值类型** -> 拷贝一份完整值，互不影响。
* **引用类型** -> 拷贝引用，共享底层数据。
* **string** 看似特殊，但在语义上仍然归入值类型。
* 复合类型里，有的属于值类型（array, struct），有的属于引用类型（slice, map, channel, interface …）。

