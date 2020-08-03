package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb) // Here the bytes shifted 10 places than kb binary
	fmt.Printf("%d\t%b\n", gb, gb)   // Here the bytes shifted 10 places than mb binary
}
