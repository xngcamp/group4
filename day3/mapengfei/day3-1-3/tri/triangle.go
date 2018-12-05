package tri

type Triangle struct {
	Heigh float64
	Wide float64
}

func (s Triangle)Area( ) float64{
	a := 0.5 * s.Heigh * s.Wide
	return a
}