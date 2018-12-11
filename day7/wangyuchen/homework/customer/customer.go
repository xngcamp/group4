package customer

import (
	"camp/day7/homework/rental"
	"fmt"
)

type Customer struct {
	name string
	rentals []*rental.Rental
}

func (c *Customer) Init(name string) {
	c.name = name
	c.rentals = make([]*rental.Rental, 0)
}

func (c *Customer) AddRental(arg *rental.Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c Customer) Statement() (res string) {
	res += fmt.Sprintf("Rental Record For %s:\n", c.name)
	cost := 0.0
	point := 0

	for i, r := range c.rentals {
		c := r.GetCost()
		res += fmt.Sprintf("\t%d.%s\t%.2f\n", i + 1, r.AMovie.GetTitle(), c)
		cost += c
		point += r.GetPoint()
	}

	res += fmt.Sprintf("Amount owned is %.2f\n", cost)
	res += fmt.Sprintf("%s earned %d frequent renter points", c.name, point)

	return
}
