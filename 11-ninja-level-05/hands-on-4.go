package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Luke Skywalker",
		friends: map[string]int{
			"Obiwan": 111,
			"Yoda":   222,
		},
		favDrinks: []string{
			"lemonade",
			"water",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.first)
	//fmt.Println(p1.friends)
	//fmt.Println(p1.favDrinks)

	for k, v := range p1.friends {
		fmt.Printf("Name: %v\n", k)
		fmt.Printf("Number: %v\n", v)
	}

	for _, v := range p1.favDrinks {
		fmt.Printf("%v\n", v)
	}
}
