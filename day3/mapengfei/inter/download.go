package main

import (
	"day3/inter/mock"
	real2 "day3/inter/real"
	"time"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download( r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main(){
	var r Retriever		// 接口里

	r = mock.Retriever{"Hello World" }			// r 里面是 mock.Retriever 对象的拷贝
	fmt.Printf("%T %v\n", r, r )
	download(r)
	inspect(r)

	r = &real2.Retriever{"www.baidu.com", time.Second * 10 }	// r 里面是 real2.Retriever 对象的指针
	fmt.Printf("%T %v\n", r, r )
	download(r)
	inspect(r)
}

func inspect(r Retriever)  {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Mock Retriever:", v.Contents)
	case *real2.Retriever:
		fmt.Println("Real Retriever:", v.UserAgent)
	}
}