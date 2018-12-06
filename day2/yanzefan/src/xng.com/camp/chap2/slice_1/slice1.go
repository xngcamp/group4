package main

import "fmt"

func main() {
	var str string
	var n int

	fmt.Println("请输入一串字符串：")
	fmt.Scanf("%s",&str)
	fmt.Println("请在0到",len(str)-1," 之间输入下标位置：")
	fmt.Scanf("%d",&n)

	fmt.Println(Slice1(str,n))
}

func Slice1(str string,j int) int {
	var num  = 0

	for i := j; i < len(str); i++ {
		if string(str[i]) == "(" {
			num++
		} else if string(str[i]) == ")" {
			num--
		}
		if num == 0 {
			return i
		}
	}
	return -1
}
