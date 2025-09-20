package visibility

import "fmt"

func PrintVar() {
	fmt.Println("PrintVar中的NyName: ", MyName)
	fmt.Println("PrintVar中的YourName: ", YourName)
	fmt.Println("PrintVar中的secret: ", secret)
}
