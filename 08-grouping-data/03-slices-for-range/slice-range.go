package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := []int{4, 5, 7, 9, 12}
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}
	// SLICE allows to group together VALUES of same TYPE
}
