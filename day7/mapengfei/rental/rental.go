package rental

import "camp/day7/mapengfei/movie"

type Rental struct {
	//	aMovie     movie.Movie
	aMovie     movie.MovieInter
	daysRented int //租期
}

func (r *Rental) Init(movie movie.MovieInter, daysRented int) {
	r.aMovie = movie
	r.daysRented = daysRented
}

func (r *Rental) GetDaysRented() int {
	return r.daysRented
}

func (r *Rental) GetMovie() movie.MovieInter {
	return r.aMovie
}

func (r Rental) GetCharge() float64 {

	return r.GetMovie().GetCharge(r.daysRented)
}

func (r Rental) GetFrequentRenterPoints() int {
	return r.GetMovie().GetFrequentRenterPoints(r.daysRented)
}
