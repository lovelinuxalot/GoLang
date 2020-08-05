package main

import "fmt"

func main() {
	x := [5]int{26, 36, 46, 56, 66}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}
