package mock



type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
//	fmt.Println("do here")
	return r.Contents
}