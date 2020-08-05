package main

import "fmt"

type persona struct {
	first_name            string
	last_name             string
	favourite_ice_flavors []string
}

func main() {
	p1 := persona{
		first_name:            "Luke",
		last_name:             "Skywalker",
		favourite_ice_flavors: []string{"vanilla", "chocolate", "peppermint"},
	}

	p2 := persona{
		first_name:            "Darth",
		last_name:             "Vader",
		favourite_ice_flavors: []string{"dark chocolate", "mango"},
	}

	fmt.Println(p1.first_name, p1.last_name)
	for _, v := range p1.favourite_ice_flavors {
		fmt.Println("\tFavorite Ice Cream Flavors: ", v)
	}
	fmt.Println(p2.first_name, p2.last_name)
	for _, v := range p2.favourite_ice_flavors {
		fmt.Println("\tFavorite Ice Cream Flavors: ", v)
	}

}
