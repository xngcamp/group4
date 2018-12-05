package rec

type Rectangle struct{
	Heigh float64
	Wide float64
}

func (s Rectangle)Area( ) float64{
	a := s.Heigh * s.Wide
	return a
}