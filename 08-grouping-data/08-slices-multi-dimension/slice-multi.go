package main

import "fmt"

func main() {
	ls := []string{"Luke", "Skywalker", "Jedi", "Green"}
	dv := []string{"Darth", "Vader", "Sith", "Red"}
	ok := []string{"Obiwan", "Kenobi", "Jedi", "Blue"}

	fmt.Println(ls)
	fmt.Println(dv)
	fmt.Println(ok)

	xp := [][]string{ls, dv, ok}
	fmt.Println(xp)

}
