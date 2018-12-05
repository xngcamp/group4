package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Println("请输入一串ISBN串：")
	fmt.Scanf("%s",&str)

	if IsValidISBN(str) == true {
		fmt.Println("这是一个正确的ISBN数")
	} else {
		fmt.Println("这不是一个正确的ISBN数")
	}

}

func IsValidISBN(str string)bool {
	var arr []int
	var temp []rune

	if str != "" {       // 判断字符串是否为空
		temp = []rune(str)
	} else {
		return false
	}
	var current []rune

	for _,x := range temp{        // 将"-"去掉
		if string(x) != "-" {
			current = append(current,x)
		}
	}

	if len(current) != 10 {
		return false
	}


	for i := 0; i < len(current) - 1 ; i++ {

		if (string(current[i]) >= "A") && (string(current[i]) <= "Z") {
			return false
		} else {
			j, err := strconv.Atoi(string(current[i]))
			if err != nil {
				panic(err)
			}
			arr = append(arr,j)
		}
	}


	if string(current[9]) == "X"  {
		arr = append(arr,10)
	} else if (string(current[9]) >= "A") && (string(current[9]) <= "Z"){
		return false
	} else {
		k, err := strconv.Atoi(string(current[9]))
		if err != nil {
			panic(err)
		}
		arr = append(arr,k)
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = arr[i] * (10 - i) + sum
	}
	if (sum % 11) == 0 {
		return true

	} else {
		return false
	}
		//return false

}
