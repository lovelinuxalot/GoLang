package main

import "fmt"

type allan int

var d allan

func main() {
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println(d)
}
