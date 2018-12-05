package main

import (
	"fmt"
	"strconv"
)

func main() {

	j, err := strconv.Atoi("X")
	if err != nil {
		panic(err)
	}
	fmt.Println(j)
}
