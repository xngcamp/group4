package map1

import "fmt"

func main()  {
	map1 := map[rune]int{} //记录字符最后出现的位置的map
	start := 0 //当前开始的位置
	max := 0 //最大值
	var arrMax  string //输出的结果
	testString:="abcabcde"
	arr:=[]rune(testString)
	for i, value := range arr {
		last := map1[value]
		if   last >= start {
			start = last + 1
		}
		if i-start+1 >max{
			if start==1{
				arrMax= string(arr[start-1:i])
			}else{
				arrMax= string(arr[start:i+1])
			}
			max=i-start+1
		}
		map1[value] = i
		//fmt.Printf("start:%d last:%d max:%d i:%d arrmax:%v\n",start,last,max,i,arrMax)
	}
	fmt.Printf("output:%v",arrMax)
}

