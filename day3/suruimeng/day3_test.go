package day3

import (
	"math"
	"testing"
)

func TestDay3(t *testing.T) {
	var test=[]struct{
		g Graph
		area float64

	}{
		{Triangle{lenght: 2,height:3},3},
		{Circle{r:1},math.Pi},
		{Rectangle{width:2,height:2},4},
	}
	for _,test:=range test{
		if get := Result(test.g);get!=test.area{
			t.Errorf("错误！")
		}
	}
}
