package day3

import (
	"fmt"
	"math"
)

type Graph interface {
	Area() float64
}

type Triangle struct {
	lenght,height float64
}

type Circle struct {
	r float64
}
type Rectangle struct {
	width,height float64
}

func (r Triangle) Area() float64 {
	return r.lenght * r.height/2
}

func (r Circle) Area() float64 {
	return math.Pi *r.r * r.r
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main()  {
	triangle :=Triangle{lenght: 2,height:3}
	circle :=Circle{r:1}
	rectangle :=Rectangle{width:2,height:2}
	Result(triangle)
	Result(circle)
	Result(rectangle)
}
<<<<<<< HEAD
func Result(g Graph)float64{
	fmt.Printf("area=%v\n",g.Area())
	return g.Area()
=======
func Result(g Graph){
	fmt.Printf("area=%v\n",g.Area())
>>>>>>> ee619d9447859f09736f4d408bfa02d43348f998
}
