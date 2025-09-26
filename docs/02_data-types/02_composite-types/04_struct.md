# 结构体 Struct

> 具体代码见：[struct结构体](../../../02_data-type/06_composite-types/04_struct)

## 定义

```go
type Person struct {
Name string
Age  int
}
```

* `type` + 名称 + `struct{}`
* 字段可以是任意类型
* **大写开头 = 导出 (public)**，可被其他包和 `json` 使用

---

## 初始化

### 零值

```go
var p Person // {Name:"", Age:0}
```

### 字面量

```go
p1 := Person{"Alice", 30}                   // 按顺序（不推荐）
p2 := Person{Name: "Bob"}                   // 命名字段（推荐）
p3 := &Person{Name: "Charlie", Age: 25} // 指针形式
p4 := new(Person)                       // *Person，零值
```

### 匿名结构体

```go
cfg := struct {
Host string
Port int
}{"localhost", 8080}
```

---

## 访问与修改

```go
p := Person{Name: "Alice"}
p.Age = 31
fmt.Println(p.Name, p.Age)
```

---

## 方法绑定

```go
func (p Person) Greet() {                // 值接收者（副本）
fmt.Println("Hi,", p.Name)
}

func (p *Person) Birthday() {            // 指针接收者（修改原值）
p.Age++
}
```

---

## 嵌入 (Embedding)

```go
type Address struct {
City string
}
type User struct {
Name string
Address // 匿名字段，方法和字段提升
}
u := User{Name: "Eve", Address: Address{City: "Paris"}}
fmt.Println(u.City) // 提升后可直接访问
```

> **注意**：Go 没有继承，嵌入是 **组合 (composition)**。

---

## 字段标签 (Tags)

```go
type Person struct {
First string `json:"firstName"`
Last  string `json:"lastName,omitempty"`
Age   int    `json:"-"` // 忽略
}
```

* `omitempty` → 零值不输出
* `"-"` → 完全忽略

---

## 指针与值的选择

* **值接收者**：只读，不修改原值，小对象常用
* **指针接收者**：可修改，避免大对象拷贝，方法接收者习惯上统一用指针

---

## 打印技巧

```go
fmt.Printf("%v\n", p) // 普通值
fmt.Printf("%+v\n", p) // 带字段名
fmt.Printf("%#v\n", p) // 带类型信息
```

---

## JSON 序列化与流式处理

```go
data, _ := json.Marshal(p) // struct → JSON
_ = json.Unmarshal(data, &p) // JSON → struct

enc := json.NewEncoder(w).Encode(p) // 写入流
dec := json.NewDecoder(r).Decode(&p) // 从流读取
```

> 只有 **导出字段** 才能被 JSON 处理。

---

## ✅ 总结

* struct = **用户自定义复合类型**，组合字段表示复杂数据。
* **大写导出**，小写仅包内可见。
* 初始化推荐用 **命名字段**。
* 行为通过 **方法接收者**附加在 struct 上。
* **嵌入 = 组合**，不是继承。
* struct tag 可定制序列化/反序列化规则。

