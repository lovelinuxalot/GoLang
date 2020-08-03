package main

import "fmt"

func main() {
	doy := 1990
	for {
		if doy > 2020 {
			break
		}
		fmt.Println(doy)
		doy++
	}
}
