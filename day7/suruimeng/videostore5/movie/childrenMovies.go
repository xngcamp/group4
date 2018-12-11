package movie

type ChildrenMovies struct {
	Title string
}

func (cm ChildrenMovies) GetTitle() string  {
	return cm.Title
}

func (cm ChildrenMovies) GetCharge(daysRented int) float64  {
	result := 0.0
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
		}
	return result
}

func (cm ChildrenMovies) GetFrequentRenterPoints(daysRented int) int  {
	return 1
}
