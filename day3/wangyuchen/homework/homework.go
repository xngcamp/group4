package main

import "fmt"

type Graph interface {
	Area() float64
}

type Triangle struct {
	Depth float64
	High float64
}

type Circle struct {
	R float64
}

type Rectangle struct {
	Depth float64
	High float64
}

func (c Circle) Area() float64{
	return c.R*c.R*3.14
}

func (t Triangle) Area() float64 {
	return t.Depth*t.High/2
}

func (r Rectangle) Area() float64 {
	return r.High*r.Depth
}

func main() {
	//var g Graph
	t := Triangle{2,3}
	fmt.Println("triangle's area is", t.Area())

	c := Circle{2}
	fmt.Println("circle's area is", c.Area())

	r := Rectangle{3,4}
	fmt.Println("rectangle's name is", r.Area())

}
