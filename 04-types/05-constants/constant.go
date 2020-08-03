package main

import "fmt"

//define constant method-1
const a = 42 //Untyped Constants
const b = 42.78
const c = "James Bond"

//define contant method-2
const (
	d int     = 42 // Typed Constants
	e float64 = 42.78
	f string  = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

}
