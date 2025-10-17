package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price,string,omitempty"`
	Discount *int    `json:"discount,omitempty"`
	Secret   string  `json:"-"`
}

// main struct 字段的 JSON tag
func main() {
	/*
		Go 的 encoding/json 包里，结构体字段的 Tag 能控制序列化/反序列化的行为
		JSON tag 必须写在导出字段（大写开头）上，否则不生效
		`json:"<名字>,<选项1>,<选项2>..."`
			<名字>：序列化后的 JSON 字段名
			<选项>：一些额外控制，比如 omitempty、string
				omitempty：如果 Age 的值是零值（0、""、false、nil），序列化时会省略这个字段 不影响反序列化
				-：完全跳过，不会序列化/反序列化
				string：序列化时强制转成字符串，反序列化时也会尝试从字符串解析

	*/
	discount1 := 20
	product1 := product{
		ID:       101,
		Name:     "LapTop",
		Price:    888.88,
		Discount: &discount1,
		Secret:   "internal-only",
	}
	jsonData1, _ := json.MarshalIndent(product1, "", "  ")
	fmt.Println(string(jsonData1))

	product2 := product{
		ID:    102,
		Name:  "Mouse",
		Price: 0,
	}
	jsonData2, _ := json.MarshalIndent(product2, "", "  ")
	fmt.Println(string(jsonData2))

}
