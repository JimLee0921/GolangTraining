package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Student struct {
	Name string
}

var studentPool = sync.Pool{
	New: func() any {
		//return new(Student)
		return &Student{Name: "default name"}
	},
}

func main() {
	/*
		由于 GC 的执行时机和 Put 的内容很难掌控
		所以通过 Get 得到数据的状态是不确定的
	*/
	// 1. 从池中获取一个对象
	s1 := studentPool.Get().(*Student)
	s1.Name = "JimLee"
	fmt.Println("first time get student from pool:", s1.Name)

	// 2. 使用完毕后放回池子
	studentPool.Put(s1)

	// 3. 再次从池中获取，可能会获取到之前放回去的对象
	s2 := studentPool.Get().(*Student)
	fmt.Println("second time get student from pool", s2.Name)

	// 3. 执行 GC
	runtime.GC()

	// 4. GC 后可能会获取到一个全新的对象
	p3 := studentPool.Get().(*Student)
	fmt.Println("after gc get student from pool", p3.Name)
}
