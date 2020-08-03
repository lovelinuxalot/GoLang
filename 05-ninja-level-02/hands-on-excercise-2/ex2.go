package main

import "fmt"

func main() {
	x := 84
	y := 65

	a := (x == y)
	b := (x <= y)
	c := (x >= y)
	d := (x != y)
	e := (x < y)
	f := (x > y)

	fmt.Println(a, b, c, d, e, f)
}
