package main

import (
	"fmt"
)

func CertfIsbn(str []string) (flag bool, des []string){
	// var isbn []string
	// isbn = append(isbn, str...)
	
	return true, str
}

func main(){
	istr := []string{"3-598-21508-8"}
	fmt.Println(CertfIsbn(istr))
}