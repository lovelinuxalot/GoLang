package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Luke",
		last:  "Skywalker",
		age:   32,
	}
	p2 := person{
		first: "Darth",
		last:  "Vader",
		age:   49,
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.age)
}
