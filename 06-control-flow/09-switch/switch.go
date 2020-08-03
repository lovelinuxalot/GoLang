package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println(" this should print 2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough //by default, finds first true case, process and exits it.
		// Use fallthrough to continue and print the next case below.
		// USe only if required.
	case (4 == 4):
		fmt.Println("Also true, but does not print?")
	default:
		fmt.Println("default case")
	}

	n := "Vader"
	switch n {
	case "Luke":
		fmt.Println("Luke Skywalker")
	case "Vader", "Anakin":
		fmt.Println("Darth Vader / Anakin Skywalker")
	case "Ren":
		fmt.Println("Kylo Ren")
	case "Solo":
		fmt.Println("Han Solo")
	default:
		fmt.Println("Sith")
	}
}
