package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func ValueOfDemo() {
	x := 10
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind()) // int
	fmt.Println(v.Int())  // 10
	// 使用 Elem 方法进行修改
	reflect.ValueOf(&x).Elem().SetInt(20)
	fmt.Println(x) // 20
}

func NewDemo() {
	// reflect.New 根据类型 t 创建一个新的指向零值的指针
	t := reflect.TypeOf(0)
	v := reflect.New(t)
	fmt.Println(v.Kind())       // ptr 指针类型
	fmt.Println(v.Elem().Int()) // 0（int 的零值）
	v.Elem().SetInt(100)        // 因为是指针，可以直接修改
	fmt.Println(v.Elem().Int()) // 100
}

func NewAtDemo() {
	x := 123
	p := unsafe.Pointer(&x)
	// 指定地址 p 强制解释为 *t
	v := reflect.NewAt(reflect.TypeOf(0), p)
	fmt.Println(v.Elem().Int()) // 123

	v.Elem().SetInt(999)

	fmt.Println(x) // 999
}

func MakeSliceDemo() {
	t := reflect.TypeOf([]int{})
	v := reflect.MakeSlice(t, 3, 5) // 长度为3，容量为5，默认都是0
	fmt.Println(v.Interface())      // [0 0 0]
	v.Index(0).SetInt(10)
	fmt.Println(v.Interface()) // [10 0 0]
}

func MakeMapDemo() {
	t := reflect.TypeOf(map[string]int{})
	v := reflect.MakeMap(t)
	v.SetMapIndex(reflect.ValueOf("A"), reflect.ValueOf(1))
	fmt.Println(v.Interface())
}

func MakeMapWithSizeDemo() {
	v := reflect.MakeMapWithSize(reflect.TypeOf(map[int]int{}), 100)
	v.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf(2))
	fmt.Println(v.Interface())
}

func MakeChanDemo() {
	t := reflect.TypeOf(make(chan int))
	ch := reflect.MakeChan(t, 2) // buffer=2
	go func() {
		ch.Send(reflect.ValueOf(10))
	}()

	v, ok := ch.Recv()
	fmt.Println(v.Int(), ok) // 10 true
}

func MakeFuncDemo() {
	adderType := reflect.TypeOf(func(int, int) int { return 0 })
	adder := reflect.MakeFunc(adderType, func(args []reflect.Value) (results []reflect.Value) {
		sum := args[0].Int() + args[1].Int()
		return []reflect.Value{reflect.ValueOf(int(sum))}
	}).Interface().(func(int, int) int)

	fmt.Println(adder(3, 7)) // 10
}

func AppendDemo() {
	s := reflect.ValueOf([]int{1, 2, 3})
	s2 := reflect.Append(s, reflect.ValueOf(4), reflect.ValueOf(5), reflect.ValueOf(6))
	fmt.Println(s2.Interface()) // [1 2 3 4 5 6]
}

func AppendSliceDemo() {
	s := reflect.ValueOf([]int{1, 2, 3})
	t := reflect.ValueOf([]int{1, 2, 3})
	// s 和 t 两个切片元素类型必须相同
	r := reflect.AppendSlice(s, t)
	fmt.Println(r.Interface()) // [1 2 3 1 2 3]
}

func IndirectDemo() {
	x := 10
	v := reflect.ValueOf(&x)

	fmt.Println(reflect.Indirect(v).Int())                  // 10
	fmt.Println(reflect.Indirect(reflect.ValueOf(x)).Int()) // 10
}

func SelectDemo() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 10
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch1)},
		{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch2)},
	}
	i, v, ok := reflect.Select(cases)
	fmt.Println(i, v.Int(), ok) // 0 10 true
}

// SliceAtDemo 有问题
func SliceAtDemo() {
	arr := [5]int{1, 2, 3, 4, 5}

	// typ 必须是切片类型：[]int
	typ := reflect.TypeOf([]int{})

	// 从 arr[2] 开始，也就是值 3 的位置
	p := unsafe.Pointer(&arr[2])

	// n = 2，表示长度和容量都是 2
	s := reflect.SliceAt(typ, p, 2)

	fmt.Println(s.Interface()) // [3 4]
}

func ZeroDemo() {
	fmt.Println(reflect.Zero(reflect.TypeOf(123)))            // 0
	fmt.Println(reflect.Zero(reflect.TypeOf("123")))          // ""
	fmt.Println(reflect.Zero(reflect.TypeOf([]int{})))        // []
	fmt.Println(reflect.Zero(reflect.TypeOf(make(chan int)))) // nil
}

func main() {
	//ValueOfDemo()
	//NewDemo()
	//NewAtDemo()
	//MakeSliceDemo()
	//MakeMapDemo()
	//MakeMapWithSizeDemo()
	//MakeChanDemo()
	//MakeFuncDemo()
	//AppendDemo()
	//AppendSliceDemo()
	//IndirectDemo()
	//SelectDemo()
	//SliceAtDemo()
	//ZeroDemo()
}
