package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := []int{4, 5, 7, 9, 12}
	fmt.Println(x)
	x = append(x, 77, 88, 1004)
	fmt.Println(x)

	y := []int{234, 546, 237}
	x = append(x, y...)
	fmt.Println(x)
}
