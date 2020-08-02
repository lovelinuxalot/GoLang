package main

import "fmt"

type hola int

var e hola

var f int

func main() {
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	e = 42
	fmt.Println(e)
	f = int(e)
	fmt.Println(f)
}
