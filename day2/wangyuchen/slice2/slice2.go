package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsValidISBN(str string) bool {
	if len(str) < 0{
		return false
	}
	str = strings.Replace(str, "-", "", -1)
	//fmt.Println(str)
	if len(str) != 10{
		return false
	}
	lastStr := str[9]
	//fmt.Println(lastStr)
	str = str[0:9]
	//fmt.Println(str)

	lastNum := 0
	//firstNum, _ := strconv.Atoi(str)

	switch lastStr {
	case 'X':lastNum = 10
	case '0':lastNum = 0
	case '1':lastNum = 1
	case '2':lastNum = 2
	case '3':lastNum = 3
	case '4':lastNum = 4
	case '5':lastNum = 5
	case '6':lastNum = 6
	case '7':lastNum = 7
	case '8':lastNum = 8
	case '9':lastNum = 9
	default:
		return false
	}
	//num := int(firstNum)*10 + lastNum
	//fmt.Println(num)
	var nums []int
	for _, v:= range str {
		i , _ := strconv.Atoi(string(v))
		nums = append(nums, i)
	}
	nums = append(nums,lastNum)
	//fmt.Printf("%v\n", nums)
	sum := 0
	i := 10
	for key,value := range nums{
		nums[key] = value
		sum = sum + nums[key]*i
		i = i-1
	}
	if sum % 11 == 0{
		return true
	}
	return false
}


func main() {
	fmt.Println(IsValidISBN("3-598-21508-8"))
}
