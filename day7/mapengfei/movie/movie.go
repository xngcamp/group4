package movie

const (
	REGULAR     = iota //0
	NEW_RELEASE        //1
	CHILDRES           // 2
)

////////////////  day7 作业  在此处定义并实现接口  /////////////////
//				在rental.go中修改定义  实现去除switch

type MovieInter interface { ///定义接口/
	GetCharge(daysRented int) float64
	GetFrequentRenterPoints(daysRented int) int
	GetTitle() string
}
type RegularMv struct {
	Title string
}
type NewMv struct {
	Title string
}
type ChildMv struct {
	Title string
}

func (r *RegularMv) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}
func (r *RegularMv) GetFrequentRenterPoints(daysRented int) int {

	return 1
}
func (r *RegularMv) GetTitle() string {
	return r.Title
}
func (n *NewMv) GetCharge(daysRented int) float64 {
	result := 0.0
	result += float64(daysRented) * float64(3)
	return result
}
func (r *NewMv) GetFrequentRenterPoints(daysRented int) int {
	result := 1
	if daysRented > 1 {
		result = 2
	}
	return result
}
func (r *NewMv) GetTitle() string {
	return r.Title
}
func (c *ChildMv) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
	}
	return result
}
func (r *ChildMv) GetFrequentRenterPoints(daysRented int) int {

	return 1
}
func (r *ChildMv) GetTitle() string {
	return r.Title
}
