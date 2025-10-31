package main

import (
	"encoding/json"
	"fmt"
)

// Telephone 模拟手机号脱敏（主要用于序列化，反序列化依旧按照原样存储）
type Telephone string

// MarshalJSON 序列化时对手机号进行脱敏处理，不暴露中间四位
func (t Telephone) MarshalJSON() ([]byte, error) {
	rawTel := string(t)
	maskTel := rawTel[:3] + "****" + rawTel[7:]
	return json.Marshal(maskTel)
}

type UserInfo struct {
	Name string    `json:"name"`
	Tel  Telephone `json:"tel"`
}

func main() {
	//u1 := UserInfo{
	//	Name: "JimLee",
	//	Tel:  "19876543210",
	//}
	//data, _ := json.Marshal(u1)
	//fmt.Println(string(data)) // {"name":"JimLee","tel":"198****3210"}

	jsonStr := `{"name":"JimLee","tel":"19876543210"}`
	var u UserInfo
	json.Unmarshal([]byte(jsonStr), &u)
	fmt.Println(u) // {JimLee 19876543210}
}
