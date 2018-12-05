package main

import (
	"fmt"
	"errors"
)

func Solution(str string, idx int)  (int, error) {
//	var stack []string	
	 var num int
	// for i := 0; i < idx; i++ {		
	// 	stack = append(stack, string(str[i]))
	// }
	// if stack[len(stack)-1] != "(" {  //判断当前位置是否为 （
	// 	return -1
	// } 	
	if str[idx] != '(' {
		return 0, errors.New("Not an opening bracket")
	}	
	for i:= idx + 1; i < len(str); i++ {
		if str[i] == '(' {
			num++
		}else if str[i] == ')'  {
			if num == 0 {
				return i, errors.New("Sucess")
			} else {
				num--
			}
		}
	}	
	
    return 0, errors.New("Not an opening bracket")
}
func main()  {
	
	str := "((1)23(45))(aB)"  //给定字符串
	n := 11  //需要索引的值
	
	fmt.Println(Solution(str, n))
	
}