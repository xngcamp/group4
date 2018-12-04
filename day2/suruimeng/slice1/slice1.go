package slice1

import "fmt"

func main()  {
	solve("((1)23(45))(aB)",2)
}
func solve(str string,index int){
	arr :=[]byte(str)
	count:=0
	for i :=index+1;i<len(arr) ;i++ {
		if arr[index]!=40{
			fmt.Printf("output:%d",-1)
			break
		}
		if arr[i]==40{
			count++
		}else if arr[i]==41{
				count--
		}
		if count==-1{
			fmt.Printf("output:%d",i)
			break
		}
	}
}
