# reflect.StructTage

StructTag 反射领域的核心模块之一，和 StructField 是绑定关系，任何结构体标签：

- `json:"name,omitempty"`
- `gorm:"column=user_name"`
- `validate:"required"`
- `yaml:"id"`

都是由这一结构支持的，主要用于

- 自定义 json 解析器
- 自定义 ORM tag 映射
- 结构体自动绑定 / 依赖注入 DI
- 表单字段映射、GraphQL、proto 属性映射

> 主要通过 StructFiled.Tag 获取

## 定义

```
type StructTag string
```

它本质就是一个字符串，但格式有强约束：

```
`key1:"value1" key2:"value2" key3:"v"`
```

## 核心方法

主要有两个核心方法：Get 和 Lookup

### Get

根据 key 取 Tag 的值，不区分该 tag 是否存在。具有局限性

```
func (tag StructTag) Get(key string) string
```

- 正常返回 string，但是 key 不存在返回 `""`
- 适合快速取值，不用判断是否存在
- 不适合区分空 tag，无法知道没写 tag 还是写了空值

### Lookup

获取 Tag 值 并可以明确知道 key 是否真实存在。

```
func (tag StructTag) Lookup(key string) (value string, ok bool)
```

- 返回值 + 是否存在，可区分空 tag 和不存在
- 框架开发必用，用于做 `JSON / ORM / YAML` 解析
- 更语义化，代码可读性 > Get()

## 语法规则

上面说 tag 不仅仅是字符串，更需要遵循特定语法：

```
`key:"value" key:"value1,value2" key2:"v"`
```

**解析行为**

- 整个 Tag 必须放在 `...` 反引号
- 每个键值格式必须是 key:"value"
- key 是 ASCII 字母 / 数字 / `-` / `_`
- 值必须包裹在双引号 `""` 中
- value 里允许含逗号用于 flags (`omitempty,string,inline`)，多 flag 以逗号分割
- 一个字段可有多个 key，但是多个 key 之间至少需要一个空格进行隔开

```
`json:name`    ❌ 不合法
`json='name'`  ❌ 必须双引号
`json:"a"db:"b"` ❌ 必须空格分隔


下面是正确示例

`json:"name"`
`json:"name" db:"user_name"`
`json:"age,omitempty,string"`
`json:"",omitempty`
`json:"-"`

`json:"id" db:"post_id"`
`json:"title" validate:"required"`
`json:"content,omitempty" xml:"content"`
```