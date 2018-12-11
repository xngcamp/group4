package rental

import "xng.com/camp/chap7/movie"

/*
Rental表示顾客租赁电影

*/

type Rental struct {
	AMovie     movie.Movie
	DaysRented int //租期
}

func (r *Rental) Init(movie movie.Movie, daysRented int) {
	r.AMovie = movie
	r.DaysRented = daysRented
}

func (r Rental) GetMovie() movie.Movie {
	return r.AMovie
}

func (r Rental) GetDaysRented() int {
	return r.DaysRented
}

func (r Rental) GetCharge() float64 {
	return r.GetMovie().GetCharge(r.DaysRented)
}

func (r Rental) GetPoints() int {
	return r.GetMovie().GetPoints(r.DaysRented)
}
