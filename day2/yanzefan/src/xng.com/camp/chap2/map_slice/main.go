package main

import (
	"fmt"
	"sort"
)

func main() {

	ns := []int{1,3}
	xs := []int{1,1,1,1,1,1}
	ys := []int{2,2,2,2,2,2}


	lastXs,lastYs := Crossover(ns,xs,ys)

	fmt.Println(lastXs,lastYs)

}

func Crossover(ns []int,xs []int,ys []int)([] int,[] int)  {

	str := make(map[int]int)
	var indexs []int

	for _,value := range ns{       // 去重复的index
		if _, ok := str[value]; !ok {
			str[value] = value
			indexs = append(indexs,value)
		}
	}
	sort.Ints(indexs)

	for i := 0; i < len(indexs) ; i++ {
		for j := indexs[i] ; j < len(xs) ; j++ {
			xs[j],ys[j] = ys[j],xs[j]
		}
	}
	return xs,ys
}
