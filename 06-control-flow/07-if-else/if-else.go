package main

import "fmt"

func main() {
	x := 40
	if x == 40 {
		fmt.Println("The value is 40")
	} else if x == 41 {
		fmt.Println("The value is 41")
	} else {
		fmt.Println("Our value is not 40")
	}
}
