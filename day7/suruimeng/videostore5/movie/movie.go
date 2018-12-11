package movie

type Movie interface {
	GetTitle() string
	GetCharge(daysRented int) float64
	GetFrequentRenterPoints(daysRented int) int
}
