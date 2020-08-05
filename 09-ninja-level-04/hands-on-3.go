package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slice1 := x[:5]
	fmt.Println(slice1)
	slice2 := x[5:]
	fmt.Println(slice2)
	slice3 := x[2:7]
	fmt.Println(slice3)
	slice4 := x[1:6]
	fmt.Println(slice4)

}
