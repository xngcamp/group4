package main

import (
		
	"fmt"
)

func LongestStr(str string) string{
//	var resMap map[string]int //存放备选结果 key为最大值 value 为结果
	begin := 0
	var result,word string
	var charMap [128]int
	for i := 0; i < len(str); i++ {
		charMap[str[i]]++
		if charMap[str[i]] == 1 {
			word  += string(str[i])
			if len(result) < len(word) {
				//resMap[result] = len(result)
			//	fmt.Println(result)
				result = word
			}
		} else {
			for ;begin < i && charMap[str[i]] > 1 ; {
				charMap[str[begin]]--
				begin++
			}
			word = ""
			for j := begin; j <= i; j++ {
				word += string(str[j])
			}
		}
	}
	// sort.Slice(resMap, func(i, j int)bool{
	// 	return  resMap[i] > resMap[j]		
	// })
//	fmt.Printf("%v",resMap)
	return result
//	return resMap[0]

}

func main(){
	test := "abcabcd"
	fmt.Println(LongestStr(test))
}