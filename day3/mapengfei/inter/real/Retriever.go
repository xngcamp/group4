package real

import (
	"os"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	Timeout time.Duration
}
func writeFile(str string){
	fileName := "E:/XNG/src/day3/inter/writetest.txt"

    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
    if err != nil {
        fmt.Println("error", err)
        os.Exit(1)
    }
    defer file.Close()

    fileString := str
    file.Seek(0, 2)    // 最后增加
    file.WriteString(fileString)
}
func (r *Retriever) Get(url string) string{
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp,true)
	fmt.Printf("%s \n", result)
	writeFile(string(result))
	return string(result)
}
