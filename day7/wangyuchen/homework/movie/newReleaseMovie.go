package movie

type NewReleaseMovie struct {
	Movie
}

func (n NewReleaseMovie) GetCost(days int) float64 {
	return float64(days) * 3
}

func (n NewReleaseMovie) GetPoint(days int) int {
	if days > 1 {
		return 2
	}
	return 1
}

func (n NewReleaseMovie) GetTitle() string {
	return n.Title
}
