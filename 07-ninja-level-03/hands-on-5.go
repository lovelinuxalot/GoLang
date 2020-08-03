package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println("When the number is %v, the modulo is %v", i, i%4)
	}
}
