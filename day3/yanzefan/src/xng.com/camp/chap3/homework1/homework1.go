package main

import "fmt"

type Graph interface {
	Area() float64
}

type Triangle struct {
	high float64
	len float64
}

type Circle struct {
	r float64
}

type Rectangle struct {
	depth float64
	len float64
}

func (triangle Triangle) Area() float64  {
	return triangle.high * triangle.len /2
}
func (circle Circle) Area() float64  {
	return 3.14 * circle.r * circle.r
}
func (rectangle Rectangle) Area() float64  {
	return rectangle.depth * rectangle.len
}

func main() {
 	var triangle Graph
	triangle = Triangle{10,20}
	var circle Graph
	circle = Circle{10}
	var rectangle Graph
	rectangle = Rectangle{10,20}
	fmt.Println(circle.Area())
	fmt.Println(triangle.Area())
	fmt.Println(rectangle.Area())

}
