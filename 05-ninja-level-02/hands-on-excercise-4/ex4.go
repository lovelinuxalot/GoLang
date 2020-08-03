package main

import "fmt"

func main() {
	a := 85
	fmt.Printf("%d\t\t%b\t\t%#x\t\t\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t\t%b\t\t%#x\t\t\n", b, b, b)

}
