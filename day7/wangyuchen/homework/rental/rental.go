package rental

import "camp/day7/homework/movie"

type Rental struct {
	AMovie     movie.GetMovie
	DaysRented int
}

func (r Rental) GetCost() float64 {
	return r.AMovie.GetCost(r.DaysRented)
}

func (r Rental) GetPoint() int {
	return r.AMovie.GetPoint(r.DaysRented)
}
