package main

import "fmt"

func main() {
	// Define struct with its values together like this
	// Use only when you need a struct in a very small scope
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Luke",
		last:  "Skywalker",
		age:   32,
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
}
