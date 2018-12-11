package movie

type NewMovie struct {
	Title string
}

func (m NewMovie) GetTitle() string {
	return m.Title
}

func (m NewMovie) GetPriceCode() int {
	return 0
}
func (m NewMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += float64(daysRented) * float64(3)
	return result
}

func (m NewMovie) GetPoints(daysRented int) int {
	if daysRented > 1 {
		return 2
	}else {
		return 1
	}
}
