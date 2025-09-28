package main

import "fmt"

// type vehicles interface{}
// Go >= 1.18 新写法，空接口没有任何方法，所以所有类型都自动实现了它
type vehicles any

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

// main 使用空接口
func main() {
	/*
		空接口
			空接口就是不包含任何方法的接口，因为接口的本质就是类型必须实现接口中的所有方法
			而空接口没有方法，所以所有类型都自动实现了它，也就是实现了多态
			在 Go 中，空接口可以表示 任意类型
		使用场景
			1. 任意类型的值（作为函数参数）
			2. 任意类型的集合（切片等）
			3. 从空接口取值（类型断言）
		Go 1.18 引入了泛型，并且官方给空接口起了一个新别名：type any = interface{}
		any 和 interface{} 是等价的，只是语义更直观：表示任意类型。

	*/

	car1 := car{vehicle: vehicle{Seats: 4, MaxSpeed: 180, Color: "Red"}, Wheels: 4, Doors: 4}
	car2 := car{vehicle: vehicle{Seats: 2, MaxSpeed: 220, Color: "Blue"}, Wheels: 4, Doors: 2}
	car3 := car{vehicle: vehicle{Seats: 5, MaxSpeed: 200, Color: "Black"}, Wheels: 4, Doors: 5}

	plane1 := plane{vehicle: vehicle{Seats: 180, MaxSpeed: 900, Color: "White"}, Jet: true}
	plane2 := plane{vehicle: vehicle{Seats: 300, MaxSpeed: 950, Color: "Silver"}, Jet: true}
	plane3 := plane{vehicle: vehicle{Seats: 50, MaxSpeed: 700, Color: "Gray"}, Jet: false}

	boat1 := boat{vehicle: vehicle{Seats: 8, MaxSpeed: 70, Color: "Blue"}, Length: 20}
	boat2 := boat{vehicle: vehicle{Seats: 12, MaxSpeed: 90, Color: "White"}, Length: 30}
	boat3 := boat{vehicle: vehicle{Seats: 4, MaxSpeed: 40, Color: "Green"}, Length: 10}
	// rides []vehicles 一个能装各种不同交通工具的集合
	riders := []vehicles{car1, car2, car3, plane1, plane2, plane3, boat1, boat2, boat3}
	// 遍历时，value 的具体类型可以是 car、plane 或 boat，但都能放进同一个切片里，因为它们都满足空接口
	for key, value := range riders {
		fmt.Println(key, " - ", value)
	}
}
