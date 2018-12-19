package movie

const (
	REGULAR     = iota //0  普通电影
	NEW_RELEASE        //1  新发行电影
	CHILDRES           //2  儿童电影
)

type Movie struct {
	Id   int
	Title     string //片名
	PriceCode int    //价格代号
	Year string      // 年份
	Author string    // 作者
	Star int        // 赞赏数
}

func (m *Movie) Init(id int,title string,priceCode int,year string,author string,star int) {
	m.Id = id
	m.PriceCode = priceCode
	m.Title = title
	m.Author = author
	m.Star = star
	m.Year = year
}

func (m Movie) GetPriceCode() int {
	return m.PriceCode
}

func (m Movie)GetTitle() string {
	return m.Title
}

func (m Movie) GetCharge(daysRented int) float64 {
	result := 0.0

	table := make(map[int]func(int,float64)(float64))

	table[REGULAR] = getRegularCharge
	table[NEW_RELEASE] = getNewReleaseCharge
	table[CHILDRES] = getChildresCharge

	return table[m.GetPriceCode()](daysRented,result)
}

func (m Movie) GetPoints(daysRented int) int {
	if m.GetPriceCode() == NEW_RELEASE && daysRented > 1 {
		return 2
	} else {
		return 1
	}
}


func getRegularCharge(daysRented int,result float64) float64 {
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}

func getNewReleaseCharge(daysRented int,result float64) float64 {
	result += float64(daysRented) * float64(3)
	return result
}

func getChildresCharge(daysRented int,result float64) float64 {
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
	}
	return result
}

