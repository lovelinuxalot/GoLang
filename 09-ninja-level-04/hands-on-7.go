package main

import "fmt"

func main() {
	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	multi_dim_slice := [][]string{slice1, slice2}

	for i, record := range multi_dim_slice {
		fmt.Println("Record: ", i)
		for i, v := range record {
			fmt.Printf("\tThe index position %v has the value %v\n", i, v)
		}
	}
}
