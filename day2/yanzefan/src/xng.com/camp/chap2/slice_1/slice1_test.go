package main

import "testing"


func TestHomework1(t *testing.T)  {

		var tests = []struct {
			str string
			n int
			i int
		}{
			{"((1)23(45))(aB)", 0,10},
			{"((1)23(45))(aB)", 1,3},
			{"((1)23(45))(aB)", 2,-1},
			{"((1)23(45))(aB)",6,9},
			{"((1)23(45))(aB)",11,14},
			{"((>)|?(*'))(yZ)",11,14},
		}

		for _, test := range tests {
			maxSize := Slice1(test.str,test.n)
			if maxSize != test.i {
				t.Errorf("Error: calMaxConLen input: %v expect: %v actual: %v ", test.str, test.n, test.i)
			}
		}





}

