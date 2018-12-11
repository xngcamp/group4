package customer

import (
	"fmt"

	"xng.com/camp/chap7/rental"
)

type Customer struct {
	name    string
	rentals []rental.Rental
}

func (c *Customer) Init(name string) {
	c.name = name
}

func (c *Customer) AddRental(arg rental.Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c Customer) GetName() string {
	return c.name
}

func (c Customer) Statement() string {
	var totalAmount float64      //消费总额
	var result = "Rental Record for " + c.GetName() + ":\n"
	for index, r := range c.rentals {

		thisAmount := r.GetCharge()

		thisResult := fmt.Sprintf("\t%d. %s\t%.2f\n", index+1, r.GetMovie().GetTitle(),thisAmount)
		result += thisResult
		totalAmount += thisAmount

	}
	// add footer lines (添加页脚信息)
	result += fmt.Sprintf("Amount owned is %.2f\n", totalAmount)
	result += fmt.Sprintf("%s earned %d frequent renter points.\n", c.GetName(), c.GetTotalPoints())
	return result

}

func (c Customer) GetTotalPoints() int {
	result := 0
	for _, r := range c.rentals {
		result += r.GetPoints()
	}
	return result
}

