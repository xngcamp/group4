package main

import (
	"camp/day7/homework/customer"
	"camp/day7/homework/movie"
	"camp/day7/homework/rental"
	"fmt"
)

func main() {
	aCustomer := new(customer.Customer)
	aCustomer.Init("Tom")

	rental1 := &rental.Rental{
		AMovie: movie.RegularMovie{
			movie.Movie{Title: "The Godfather", PriceCode: movie.REGULAR},
		},
		DaysRented: 5,
	}
	aCustomer.AddRental(rental1)

	rental2 := &rental.Rental{
		AMovie: movie.NewReleaseMovie{
			movie.Movie{Title: "Fast & Furious", PriceCode: movie.NEW_RELEASE},
		},
		DaysRented: 1,
	}
	aCustomer.AddRental(rental2)

	rental3 := &rental.Rental{
		AMovie: movie.ChildrenMovie{
			movie.Movie{Title: "Casablanca", PriceCode: movie.CHILDREN},
		},
		DaysRented: 7,
	}
	aCustomer.AddRental(rental3)

	result := aCustomer.Statement()
	fmt.Println(result)
}
