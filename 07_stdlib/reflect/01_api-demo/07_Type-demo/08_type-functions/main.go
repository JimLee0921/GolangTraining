package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}
type B interface {
}

func TypeOfDemo() {
	fmt.Println(reflect.TypeOf("Hello")) // string
	fmt.Println(reflect.TypeOf([]int{})) // []int
	fmt.Println(reflect.TypeOf(&User{})) // *main.User
	fmt.Println(reflect.TypeOf(User{}))  // main.User
	fmt.Println(reflect.TypeOf(nil))     // <nil>
}

func TypeForDemo() {
	t := reflect.TypeFor[User]()
	fmt.Println(t) // main.User
	t2 := reflect.TypeFor[*User]()
	fmt.Println(t2) // main.*User
}

func ArrayOfDemo() {
	t := reflect.ArrayOf(5, reflect.TypeOf(0))
	fmt.Println(t) // [5]int
	t2 := reflect.ArrayOf(4, reflect.TypeOf("JimLee"))
	fmt.Println(t2) // [4]string
}

func SliceOfDemo() {
	t := reflect.SliceOf(reflect.TypeOf("haha"))
	fmt.Println(t) // []string
	t2 := reflect.SliceOf(reflect.SliceOf(reflect.TypeOf("enen")))
	fmt.Println(t2) // [][]string
}

func PointerToDemo() {
	t := reflect.PointerTo(reflect.TypeOf(0))
	fmt.Println(t) // *int
	t2 := reflect.PointerTo(reflect.TypeOf(User{}))
	fmt.Println(t2) // *main.User
}

func MapOfDemo() {
	t := reflect.MapOf(
		reflect.TypeOf(""),
		reflect.TypeOf(0),
	)
	fmt.Println(t) // map[string]int
	t2 := reflect.MapOf(
		reflect.TypeOf(0),
		reflect.SliceOf(reflect.TypeOf(byte(0))),
	)
	fmt.Println(t2) // map[int][]uint8
}

func ChanOfDemo() {
	t := reflect.ChanOf(reflect.BothDir, reflect.TypeOf(0))
	fmt.Println(t) // chan int 双向通道
	t2 := reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(""))
	fmt.Println(t2) // <-chan string 只读
	t3 := reflect.ChanOf(reflect.SendDir, reflect.TypeOf(User{}))
	fmt.Println(t3) // chan<- main.User 只写
}

func FuncOfDemo() {
	in := []reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")}
	out := []reflect.Type{reflect.TypeOf(true)}

	t := reflect.FuncOf(in, out, false)
	fmt.Println(t) // func(int, string) bool

	// 可变参数
	t2 := reflect.FuncOf([]reflect.Type{reflect.TypeOf([]int{})}, nil, true)
	fmt.Println(t2) // func(...int)
}

func StructOfDemo() {
	fields := []reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
			Tag:  `json:"name"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
			Tag:  `json:"age"`,
		},
	}

	t := reflect.StructOf(fields)
	fmt.Println(t) // struct { Name string "json:\"name\""; Age int "json:\"age\"" }
	// 可以创建实例
	v := reflect.New(t).Elem()
	v.FieldByName("Name").SetString("Jimmy")
	v.FieldByName("Age").SetInt(20)

	fmt.Println(v.Interface()) // {Jimmy 20}
}

func main() {
	//TypeOfDemo()
	//TypeForDemo()
	//ArrayOfDemo()
	//SliceOfDemo()
	//PointerToDemo()
	//MapOfDemo()
	//ChanOfDemo()
	//FuncOfDemo()
	StructOfDemo()
}
