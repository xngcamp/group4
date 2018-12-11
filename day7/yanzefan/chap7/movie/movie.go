package movie

type Movie interface {

	GetTitle() string
	GetPriceCode() int
	GetCharge(daysRented int) float64
	GetPoints(daysRented int) int
}
