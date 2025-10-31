package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name   string `json:"user_name"`
	Age    int    `json:"age"`
	Secret string `json:"-"`                  // 不参与编解码
	Money  int    `json:"my_money,omitempty"` // 编码时零值不显示
}

func DecoderWithTagDemo() {
	jsonStr := `{"user_name":"JimLee","age":20, "Secret": "My Secret", "my_money":0}`
	r := strings.NewReader(jsonStr)

	dec := json.NewDecoder(r)

	var u User
	if err := dec.Decode(&u); err != nil {
		panic(err)
	}

	fmt.Println(u) // {JimLee 20  0}
}

func EncodeWithTagDemo() {
	u1 := User{
		Name:   "JimLee",
		Age:    20,
		Secret: "My Secret",
		Money:  2,
	}
	u2 := User{
		Name:   "FrankStan",
		Age:    15,
		Secret: "My Secret",
		Money:  0,
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(u1); err != nil { // {"user_name":"JimLee","age":20,"my_money":2}
		panic(err)
	}

	if err := enc.Encode(u2); err != nil { // {"user_name":"FrankStan","age":15}
		panic(err)
	}
}

func main() {
	DecoderWithTagDemo()
	EncodeWithTagDemo()
}
