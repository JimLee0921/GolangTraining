# strings 包

strings 是 Go 标准库中专门用来操作字符串的工具包。

> 文档地址：https://pkg.go.dev/strings
> 
> https://go.dev/blog/strings

## 主要功能

Go 里的字符串类型是不可变的（immutable），也就是说：

```
s := "hello"
s[0] = 'H' // 这样是不能改的
```

所以当要进行：

- 查找内容
- 替换内容
- 分割、拼接
- 大小写切换
- 去空格
- 检查前缀 / 后缀

这些字符串处理操作时，就需要依赖 `strings` 包。

## 局限性

strings 包只处理 `ASCII/UTF-8` 字符串，不会处理 Unicode 的字符宽度问题，比如中文、emoji 长度问题，需要使用 `unicode/utf8`
包解决

