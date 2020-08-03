package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s) // String

	bs := []byte(s) // Convert to Bytes size
	fmt.Println(bs)

	n := bs[0] // Checking first value of the byte
	fmt.Println(n)
	fmt.Printf("%T\n", n)  // TYPE of the byte == uint8
	fmt.Printf("%b\n", n)  // Binary
	fmt.Printf("%x\n", n)  // Hexadecimal
	fmt.Printf("%X\n", n)  // Hexadecimal with caps
	fmt.Printf("%#x\n", n) // Hexadecimal with notation

}
