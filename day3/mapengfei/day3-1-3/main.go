package main

import (
	"fmt"
	"day3/day3-1-3/rec"
	"day3/day3-1-3/cir"
	"day3/day3-1-3/tri"
)

type Graph interface {
	Area() float64
}

func CalcuArea(s Graph) float64{
	return s.Area()
}

func main(){
	r := rec.Rectangle {
		Heigh: 2,
		Wide: 3,
	}	
	fmt.Printf("长方形的面积：%f \n",CalcuArea(r))
	
	t := tri.Triangle {
		Heigh: 2,
		Wide: 3,
	}
	fmt.Printf("三角形的面积：%f \n",CalcuArea(t))

	c := cir.Circle {
		R: 2,
	}
	fmt.Printf("圆形的面积：%f \n",CalcuArea(c))

	
}