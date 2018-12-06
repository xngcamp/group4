package main

import (
	"testing"
)


func TestCrossover(t *testing.T)  {

		var tests = []struct {
			ns []int
			xs []int
			ys []int
			lastXs []int
			lastYs []int
		}{
			{[]int{2}, []int{1,1,1,1,1,1},[]int{2,2,2,2,2,2},
				[]int{1,1,2,2,2,2},[]int{2,2,1,1,1,1}},
			{[]int{1,3}, []int{1,1,1,1,1,1},[]int{2,2,2,2,2,2},
				[]int{1,2,2,1,1,1},[]int{2,1,1,2,2,2}},
		}

		for _, test := range tests {
			tempXs,tempYs := Crossover(test.ns,test.xs,test.ys)
			if !equal(tempXs,test.lastXs) && !equal(tempYs,test.lastYs){
				t.Errorf("Error: Crossover input: %v,%v,%v expect: %v,%v actual: %v,%v",
					test.ns,test.xs,test.ys, test.lastXs, test.lastYs,tempXs,tempYs)
			}
		}

}

func equal(xs []int, ys []int) bool {
	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {

			return false
		}
	}
	return true
}

