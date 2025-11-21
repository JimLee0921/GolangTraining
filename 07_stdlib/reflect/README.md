# reflect 模块

reflect 是 Go 语言提供的 运行时反射（runtime reflection） 的标准库，它让程序能够在运行时检查、修改变量的类型和值。


> 文档地址：https://pkg.go.dev/reflect

## 主要作用

- 动态查看变量类型：例如不知道变量是什么类型，可以用反射判断它是 int, struct, slice 等
- 动态获取或修改变量的值：例如把一个变量从 10 改成 20，或者读取 struct 字段的内容
- 动态遍历结构体字段：比如 ORM 框架（GORM）就是用 reflect 遍历字段、tag、类型等实现结构体映射到数据库的
- 调用未知方法：比如可以在运行时调用 obj.MethodByName("Save")() 这样的动态方法

## 学习顺序

1. 先理解两个核心概念：
    - Type（类型信息）
    - Value（值信息）
2. 学会如何判断类型（Kind）
3. 学会取值 / 改值（需要指针）
4. 学会操作 Struct（读取字段、tag）
5. 学会调用方法（MethodByName）