package movie

const (
	REGULAR = iota
	NEW_RELEASE
	CHILDREN
)

type GetMovie interface {
	GetCost(days int) float64
	GetPoint(days int) int
	GetTitle() string
}

type Movie struct {
	Title     string
	PriceCode int
}