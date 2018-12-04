package isbn

func IsValidISBN(str string)bool  {
	arr :=[]byte(str)
	for i :=0;i<len(arr) ;i++ {
		if arr[i]==45{
			arr=remove(arr,i)
		}
	}
	arrInt :=byteToInt(arr)
	count:=0
	for i ,j:=0,len(arrInt);j>0 && i<len(arrInt) ;i,j=i+1,j-1 {
		count+=arrInt[i]*j
	}
	if count%11==0{
		return true
	}else {
		return false
	}

}

func byteToInt(by []byte)[]int  {
	var result []int
	for i:=range by{
		switch by[i] {
		case 48:
			result= append(result, 0)
		case 49:
			result= append(result, 1)
		case 50:
			result= append(result, 2)
		case 51:
			result= append(result, 3)
		case 52:
			result= append(result, 4)
		case 53:
			result= append(result, 5)
		case 54:
			result= append(result, 6)
		case 55:
			result= append(result, 7)
		case 56:
			result= append(result, 8)
		case 57:
			result= append(result, 9)
		case 88:
			result= append(result, 10)
		}
	}
	return result
}

func remove(slice []byte,index int)[]byte  {
	copy(slice[index:],slice[index+1:])
	return slice[:len(slice)-1]
}
