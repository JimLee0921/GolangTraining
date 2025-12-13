# 示例测试 (Example Test)

Example Test 示例测试是 Go testing 体系中一种特殊测试形式,通过对比输出文本来判断是否通过测试，主要可以做：

- 可执行的测试
- 可验证的文档示例
- `pkg.go.dev` 自动展示的示例代码来源
- 使用 `go test` 时Example 测试是默认测试流程的一部分，也可以使用 `go test -run Example` 只运行示例测试

主要生成代码文档和验证代码逻辑，表示一段代码如何使用并且用法是否还是正确的

## 函数基本格式

这些格式不是建议，而是识别规则，必须遵守

```
func Example()
func ExampleFunc()
func ExampleType()
func ExampleType_Method()
func Example<Identifier>_<Suffix>
```

### Example

```
func Example() {
    // ...
}
```

- package 包级别示例
- 不绑定任何具体标识符
- 在 `pkg.go.dev` 中显示在 package 示例区域

### ExampleFunc

```
func ExampleParse() {
    // ...
}
```

- 绑定到函数 Parse
- `pkg.go.dev` 会把它展示在 `Parse` 的文档下

### ExampleType

```
func ExampleClient() {
    // ...
}
```

- 绑定到类型 `Client` 上
- 展示在该类型的文档中

### ExampleType_Method

```
func ExampleClient_Do() {
    // ...
}
```

- 绑定到类型 `Client` 的 `Do` 方法上
- 这是方法级别 Example 的唯一合法命名形式

### Example<Identifier>_<Suffix>

创建子 Example

- Identifier：必须是某个已经存在的标识符：函数名/类型名/类型的函数名（Type_Method）
- Suffix：任意非空标识字符串，官方与社区约定使用小写

> Go 不关系 Suffix 是什么，全不解析语义，只当作这个标识符的另一个示例

```
func ExampleParse_error() {}
func ExampleParse_error_timeout() {}
func ExampleParse_ok() {}
func ExampleParse_invalidInput() {}
func ExampleParse_withCache() {}
func ExampleParse_case1() {}
func ExampleParse_中文() {}      // 技术上可行，但强烈不推荐
```

在 `pkg.go.dev` 中`_suffix` 会去掉下划线，放在括号里，作为子标题，展示效果为：

```
Parse
    Example
    Example (error)
    Example (error_timeout)
    ...
```

> 注意这种 Example 的写法与 Type_Method 名字并不冲突，Go 在解析 Example 时：
> 1. 从 Example 后面开始
> 2. 贪婪匹配一个合法的标识符（函数/类型/方法）
> 3. 只要在匹配成功后才会把剩余部分作为 suffix

## Output 示例验证

在 `Example` 测试中，是否有 `//Output` 是执行与否的分界线

- 有 `//Output`：`Example` 会被执行，`Output` 中的输出内容会被校验
- 没有 `//Output`：`Example` 只会进行编译，不会执行，不参与测试成败

### Output 比较方式

在 Example 测试中：

- 比对的是标准输出 `stdout`
- 如果有多个运行结果则按行进行比较
- 内容/顺序/换行 都必须保持一致才能测试通过
- 不要使用 `time.Now()` 这种非确定性输出，`Output`内容不支持正则
- 可以使用 `// Unordered output`无序输出匹配，主要用于 map 遍历，不知道输出顺序的场景

```
func ExampleAdd() {
    fmt.Println(Add(1, 2))
    // Output:
    // 3 
}
```

## 文档生成

运行 `go doc -all your/module/path` 或生成在线文档时（如 pkg.go.dev），示例函数会自动出现在文档里，变成可运行的示例代码

## 注意事项

| 限制                         | 说明                      |
|----------------------------|-------------------------|
| 函数不能有参数                    | 必须是 `func ExampleXxx()` |
| 无法使用 `t *testing.T`        | 不能调用 `t.Error()` 等函数    |
| 输出严格匹配                     | 空格、换行都要一致               |
| 默认比较标准输出                   | 不比较返回值或日志               |
| 可以用 `// Unordered output:` | 表示输出顺序不要求一致（Go 1.15+）   |

