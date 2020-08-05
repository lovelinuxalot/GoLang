package main

import "fmt"

type person struct {
	first_name            string
	last_name             string
	favourite_ice_flavors []string
}

func main() {
	p1 := person{
		first_name:            "Luke",
		last_name:             "Skywalker",
		favourite_ice_flavors: []string{"vanilla", "chocolate", "peppermint"},
	}

	p2 := person{
		first_name:            "Darth",
		last_name:             "Vader",
		favourite_ice_flavors: []string{"dark chocolate", "mango"},
	}

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println("\t", v.first_name)
		fmt.Println("\t", v.last_name)
		fmt.Println("\t Favourite Ice Cream Flavours:")
		for i, val := range v.favourite_ice_flavors {
			fmt.Printf("\t\t%v %v\n", i+1, val)
		}
		fmt.Println("---------------")
	}

}
