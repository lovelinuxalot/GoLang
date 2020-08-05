package main

import "fmt"

func main() {
	// x := type{values} // COMPOSITE LITERAL
	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[1] = 23
	x[9] = 56
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//x[10] = 12  // THis does not work. Index out of range. to do that we need to append
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
