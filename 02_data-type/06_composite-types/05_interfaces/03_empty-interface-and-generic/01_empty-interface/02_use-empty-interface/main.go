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
		Go 1.18 后并且官方给空接口起了一个新别名：type any = interface{}
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
