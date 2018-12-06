package main

import "fmt"

func main() {

	str := "abcdefbgac"
	FindLongest(str)

}

func FindLongest( s string )  {

	lastOccurred := make(map[byte]int)

	var arr []byte

	start := 0
	maxLen := 0

	for i, ch := range []byte(s) {

		lastI, ok := lastOccurred[ch]

		if ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}

		if i - start + 1 > maxLen {
			arr = append(arr,ch)
			maxLen = i - start + 1
		}

		lastOccurred[ch] = i

	}

	fmt.Printf("%q",arr)

}
