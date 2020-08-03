package main

import "fmt"

func main() {
	favSport := "swimming"
	switch favSport {
	case "skiing":
		fmt.Println("Go to the Mountains")
	case "surfing":
		fmt.Println("Go to Hawaii")
	case "swimming":
		fmt.Println("Go to a pool")
	case "wingsuit flying":
		fmt.Println("What would you like me to say at your funeral?")
	}
}
