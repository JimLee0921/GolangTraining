package main

import (
	"fmt"
	"reflect"
)

type Base struct {
	BaseID int `json:"id"`
}

type User struct {
	Base
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age,omitempty"`
	secret string // 私有字段，不导出
}

//
//// 定义函数将任意 struct 转为 map，模仿 JSON Tag 进行解析
//func ConvertStruct(s any) map[string]any {
//	result := map[string]any{}
//	v := reflect.ValueOf(s)
//	t := reflect.TypeOf(s)
//
//	// 必须是 struct
//	if t.Kind() != reflect.Struct {
//		panic("ConvertStrict only accept struct")
//	}
//
//	// 遍历所有字段
//
//}

func TypeDemo() {
	// 使用 reflect.TypeOf 获取 StructField
	t := reflect.TypeOf(User{})
	fmt.Println(t.NumField()) // Fields 数量
	// 下标配合 t.NumField 遍历所有字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Type)
		fmt.Println(f.Name)
		fmt.Println(f.Tag)
		fmt.Println(f.Index)
		fmt.Println(f.Anonymous)
		fmt.Println(f.Offset)
		fmt.Println(f.PkgPath)
	}
	// 按照字段名查找
	f, ok := t.FieldByName("Age")
	if ok {
		fmt.Println(f.Name, f.Type) // Age int
	}

	// FieldByIndex([]int) 用于匿名/嵌套字段
	nestField, _ := t.FieldByName("BaseID")
	fmt.Println(nestField.Index)                      // [0 0]
	fmt.Println(t.FieldByIndex(nestField.Index).Name) // BaseID
}

func ValueDemo() {
	// 使用 reflect.ValueOf 获取字段真实值
	u := User{
		Base: Base{
			BaseID: 6,
		},
		ID:     1,
		Name:   "JimLee",
		Age:    20,
		secret: "My secret",
	}
	v := reflect.ValueOf(&u).Elem() // 传入指针，获取字段真实值，可以用于修改
	val1 := v.Field(0)              // 下标获取
	val2 := v.FieldByName("Age")    // 字段名获取
	fmt.Println(val1, val2)         // {6} 20
	// 修改值
	val1.Set(reflect.ValueOf(Base{BaseID: 123}))
	val2.SetInt(66)
	fmt.Println(val1, val2) // {123} 66

}

func main() {
	//TypeDemo()
	ValueDemo()
}
