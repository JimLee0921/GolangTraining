package main

import "fmt"

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

func main() {
	/*
		这里定义三个嵌入了 vehicle 的结构体：car，plane 和 boat
		但是在使用时还是遍历打印还是需要使用不同的类型切片
	*/
	car1 := car{vehicle: vehicle{Seats: 4, MaxSpeed: 180, Color: "Red"}, Wheels: 4, Doors: 4}
	car2 := car{vehicle: vehicle{Seats: 2, MaxSpeed: 220, Color: "Blue"}, Wheels: 4, Doors: 2}
	car3 := car{vehicle: vehicle{Seats: 5, MaxSpeed: 200, Color: "Black"}, Wheels: 4, Doors: 5}
	cars := []car{car1, car2, car3}

	plane1 := plane{vehicle: vehicle{Seats: 180, MaxSpeed: 900, Color: "White"}, Jet: true}
	plane2 := plane{vehicle: vehicle{Seats: 300, MaxSpeed: 950, Color: "Silver"}, Jet: true}
	plane3 := plane{vehicle: vehicle{Seats: 50, MaxSpeed: 700, Color: "Gray"}, Jet: false}
	planes := []plane{plane1, plane2, plane3}

	boat1 := boat{vehicle: vehicle{Seats: 8, MaxSpeed: 70, Color: "Blue"}, Length: 20}
	boat2 := boat{vehicle: vehicle{Seats: 12, MaxSpeed: 90, Color: "White"}, Length: 30}
	boat3 := boat{vehicle: vehicle{Seats: 4, MaxSpeed: 40, Color: "Green"}, Length: 10}
	boats := []boat{boat1, boat2, boat3}

	for key, value := range cars {
		fmt.Println(key, " - ", value)
	}

	for key, value := range planes {
		fmt.Println(key, " - ", value)
	}

	for key, value := range boats {
		fmt.Println(key, " - ", value)
	}
}
