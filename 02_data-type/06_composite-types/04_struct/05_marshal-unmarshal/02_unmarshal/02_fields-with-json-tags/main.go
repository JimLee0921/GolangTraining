package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price,string,omitempty"`
	Discount *int    `json:"discount,omitempty"`
	Secret   string  `json:"-"`
}

func main() {
	/*
		缺少字段:
			price 缺省 → 赋零值 0.0
			discount 缺省 → 保持 nil
			omitempty 只影响 序列化，对反序列化没影响
		多余字段
			Go 的 json.Unmarshal 是宽容模式
			有用的就解析，没定义的字段就丢掉
			有多余的字段是反序列化不会报错
	*/

	// 模拟 JSON 数据
	raw1 := `{"id":101,"name":"Laptop","price":"899.99","discount":20}`
	var p1 Product
	_ = json.Unmarshal([]byte(raw1), &p1)

	fmt.Printf("%+v\n", p1)
	// 缺少字段
	raw2 := `{"id":102,"name":"Mouse"}`
	var p2 Product
	_ = json.Unmarshal([]byte(raw2), &p2)
	fmt.Printf("%+v\n", p2)

	// 多字段 extra 会被直接忽略，反序列化不会报错
	raw3 := `{"id":103,"name":"Keyboard","extra":"ignore me"}`
	var p3 Product
	_ = json.Unmarshal([]byte(raw3), &p3)
	fmt.Printf("%+v\n", p3)

}
