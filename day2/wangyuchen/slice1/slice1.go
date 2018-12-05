package main

import "fmt"

func Slice1(str string, idx int) int {
	count := 0
	for i := idx; i < len(str); i++ {
		if str[i] == '('{
			count ++
		}
		if str[i] == ')'{
			count --
			if count == 0{
				return i
			}
		}
	}
	if count != 0{
		return -1
	}
	return 0
}

func main() {
	fmt.Println(Slice1("((1)23(45))(aB)", 6))
	//fmt.Println(Slice1("((>)|?(*'))(yZ)", 11))
}
