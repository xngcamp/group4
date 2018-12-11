package movie

type RegularMovie struct {
	Title string
}

func (m RegularMovie) GetTitle() string {
	return m.Title
}

func (m RegularMovie) GetPriceCode() int {
	return 0
}
func (m RegularMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}

func (m RegularMovie) GetPoints(daysRented int) int {
	return 1
}