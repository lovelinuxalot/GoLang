package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := []int{4, 5, 7, 9, 12}
	fmt.Println(x)
	fmt.Println(x[1:])  // From index 1 to end of slice
	fmt.Println(x[1:3]) // From index 1 to index 3, where the first value is included and the last value is excluded
	fmt.Println(x[:1])  // From index start to index 1, but not including index 1
	fmt.Println(x[:2])  // From index start to index 2, but not including index 2

	// FOR loop to print the values with index. Both methods work
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
