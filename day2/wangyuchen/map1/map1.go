package main

import "fmt"

func Map1(str string) string{
	last_location := make(map[rune]int)
	start := 0
	max := 0
	res := ""
	for key, value := range str{
		if last, exists := last_location[value]; exists && last >= start {
			start = last + 1
		}
		if key-start+1> max{
			max = key- start +1
			res = str[start:key+1]
		}
		last_location[value] = key
	}
	return res
}


func main (){
	//fmt.Println(Map1("abababa"))
	fmt.Println(Map1("cdefbga"))
}
