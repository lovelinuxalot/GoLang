package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 20 {
			fmt.Println("Got in if condition and breaking!!!")
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("program exited. done")
}
