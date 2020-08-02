package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)                // Type
	fmt.Printf("%v\n", y)                // Value in default format
	fmt.Printf("%b\n", y)                // Binary
	fmt.Printf("%x\n", y)                // Hexadecimal
	fmt.Printf("%#x\n", y)               // Hexadecimal with 0x
	fmt.Printf("%#x\t%b\t%x\t", y, y, y) //Printing combination of values

	s := fmt.Sprintf("%#x\t%b\t%x\t", y, y, y) //Printing as String
	fmt.Println(s)
}
