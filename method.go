package main

import "fmt"

func main() {
	rec1 := Rectangle{
		length: 2,
		width:  4,
	}
	fmt.Println("Area Before:", rec1.Area())
	rec1.Scale(1.5)
	fmt.Println("length:", rec1.length)
	fmt.Println("width:", rec1.width)
	fmt.Println("Area After:", rec1.Area())

	myNumber := myInt(-5)
	fmt.Println(myNumber.isPositive())
	fmt.Println(myNumber.Welcom())

	s := Shape{Rectangle: Rectangle{length: 10, width: 2}}
	fmt.Println("Shape Area:", s.Area())

}

type Rectangle struct {
	length float64
	width  float64
}

// value reciver method
// We use a value receiver if the method does not modify the receiver instance.
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// we use a pointer receiver if the method needs to modify the receiver instance, or if you want to avoid copying large structs.
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type myInt int

// we can associate methods on any type that we want.
func (m myInt) isPositive() bool {
	return m > 0
}

// there is no need to add an instance if not using it, you can just assing a method to a type
func (myInt) Welcom() string {
	return "welcome to my type"
}

// promotion
type Shape struct {
	Rectangle
}
