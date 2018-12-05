package main

import "fmt"

func main() {

	str := "ababab"

	var temp []rune
	for _ , x := range str{
		temp = append(temp,x)
	}

	fmt.Println(temp)

	var arr map[string]int

	for i := 0; i < len(temp); i++ {
		for j := i + 1; j < len(temp); j++ {
			if temp[i] {

			}
		}
	}

}
