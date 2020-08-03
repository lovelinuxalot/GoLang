package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is true")
	}
	if false {
		fmt.Println("This is false")
	}
	if !true {
		fmt.Println("This is not true")
	}
	if !false {
		fmt.Println("This is not false")
	}
	if 2 == 2 {
		fmt.Println("This is true expression")
	}
	if 2 != 2 {
		fmt.Println("This is false expression")
	}
	if !(2 == 2) {
		fmt.Println("This is not true expression")
	}
	if !(2 != 2) {
		fmt.Println("This is not false expression")
	}

}
