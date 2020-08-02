package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) //prints FALSE, as ZERO VALUE is ASSIGNED without VALUE
	x = true
	fmt.Println(x) //prints TRUE, VALUE ASSIGNED
}
