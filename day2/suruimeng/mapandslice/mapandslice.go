package main

import (
	"fmt"
	"sort"
)

func main()  {
	Crossover([]int{1,3},[]int{1,2,3},[]int{10,11,12})
}

func Crossover(ns [] int,xs [] int,ys [] int)  {
	var map1= map[string][]int{
		"xs":[]int{},
		"ys":[]int{},
	}
	sort.Ints(ns)
	if len(xs)!=len(ys)||ns[len(ns)-1]>len(xs){
		fmt.Errorf("传入参数错误")
		return
	}
	for i:=range ns{
		map1["xs"]=xs
		map1["ys"]=ys
		xs=append(map1["xs"][:ns[i]],map1["ys"][ns[i]:]...)
		ys=append(map1["ys"][:ns[i]],map1["xs"][ns[i]:]...)
		map1["xs"]=xs
		map1["ys"]=ys
	}
	fmt.Printf("xs:%v,ys:%v \n",map1["xs"],map1["xs"])
}
