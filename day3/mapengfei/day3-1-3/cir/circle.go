package cir

type Circle struct {
	R float64
}

func (s Circle)Area( ) float64{
	a := 3.141592 * s.R * s.R
	return a
}