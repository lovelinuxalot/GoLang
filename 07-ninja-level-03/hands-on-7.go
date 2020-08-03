package main

import "fmt"

func main() {
	x := "MoneyPenny"
	if x == "James Bond" {
		fmt.Println(x)
	} else if x == "MoneyPenny" {
		fmt.Println("MoneyPenny")
	} else {
		fmt.Println("Neither of both")
	}
}
