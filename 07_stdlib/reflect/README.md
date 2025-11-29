# reflect 模块

reflect 是 Go 语言提供的 运行时反射（runtime reflection） 的标准库，它让程序能够在运行时检查、修改变量的类型和值。


> 文档地址：https://pkg.go.dev/reflect

## 主要作用

- 动态查看变量类型：例如不知道变量是什么类型，可以用反射判断它是 int, struct, slice 等
- 动态获取或修改变量的值：例如把一个变量从 10 改成 20，或者读取 struct 字段的内容
- 动态遍历结构体字段：比如 ORM 框架（GORM）就是用 reflect 遍历字段、tag、类型等实现结构体映射到数据库的
- 调用未知方法：比如可以在运行时调用 obj.MethodByName("Save")() 这样的动态方法

## 主要核心概念

Type 问是什么，Value 管干什么

| 对象              | 主要能力                          | 
|-----------------|-------------------------------|
| `reflect.Type`  | 查看信息：字段、方法、容器结构、Kind、Size、Tag |
| `reflect.Value` | 操作值本身：取值、设值、调用方法、遍历容器         |


> 获取 Type 的主要入口是 `reflect.TypeOf` 和新版本的 `reflect.TypeFor`
> 获取 Value 的主要入口是 `reflect.ValueOf`