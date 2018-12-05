package main

import "testing"

func TestCalcuArea(t *testing.T) {
	var data = []struct{
		g Graph
	//	area float64
		want float64
	}{
		{Triangle{Height: 2, Wide: 2}, 4.000000 },
	}

	for _, s := range data {
		got := CalcuArea(s.g)
		if got != s.want {
			t.Error("计算错误")
		} else {
			t.Error("计算正确")
		}
	}

}