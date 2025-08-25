package basic

import "fmt"

const PI = 3.14

func main() {
	fmt.Println("Hello " + "World")

	// var age int
	name := "John"

	const (
		saturday = 1
		sunday   = 2
		monday   = 3
	)
	const NUMBER int = 12

	fmt.Println("name is: ", name)
	fmt.Println("const value is: ", PI)
}
