package main

import (
	"fmt"
)

type Shape interface {
	area() float32
	perimeter() float32
}

type Rectangle struct {
	l float32
	b float32
}

type Square struct {
	a float32
}

func (rect Rectangle) area() (area float32) {
	area = rect.l * rect.b
	return
}
func (rect *Rectangle) perimeter() (perimeter float32) {
	perimeter = 2 * (rect.l + rect.b)
	return
}
func (sqr *Square) area() (area float32) {
	area = sqr.a
	return

}
func (sqr *Square) perimeter() (perimeter float32) {
	perimeter = 4 * sqr.a
	return
}

func main() {

	var shape Shape
	rect := Rectangle{}
	sqr := Square{}

	fmt.Println("Enter rectangle length:")
	fmt.Scan(&rect.l)
	fmt.Println("Enter rectangle bredth:")
	fmt.Scan(&rect.b)
	fmt.Println("Enter square length:")
	fmt.Scan(&sqr.a)

	shape = &rect
	area := shape.area()
	fmt.Println("Area:", area)
	perimeter := shape.perimeter()
	fmt.Println("Perimeter:", perimeter)

	shape = &sqr
	area = shape.area()
	fmt.Println("Area:", area)
	perimeter = shape.perimeter()
	fmt.Println("Perimeter:", perimeter)

}
