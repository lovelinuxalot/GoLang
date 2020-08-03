package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x) //or
	fmt.Printf("%#x\n", x)
	//OR in one line
	fmt.Printf("%d\t\t%b\t\t%#x\t\t", x, x, x)
}
