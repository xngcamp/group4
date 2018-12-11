package movie

type NewResleaseMovies struct {
	Title string
}

func (nrm NewResleaseMovies) GetTitle() string  {
	return nrm.Title
}

func (nrm NewResleaseMovies) GetCharge(daysRented int) float64  {
	result := 0.0
	result += float64(daysRented) * float64(3)
	return result
}

func (nrm NewResleaseMovies) GetFrequentRenterPoints(daysRented int) int  {
	if daysRented>1{
		return 2
	}
	return 1
}
