package main

import "fmt"

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	// Another way to do with for
	for {
		if x > 9 {
			fmt.Println("Got in if condition")
			break
		}
	}
	fmt.Println("program exited. done")
}
