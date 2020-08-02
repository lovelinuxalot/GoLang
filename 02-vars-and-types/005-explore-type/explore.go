package main

import "fmt"

var y = 42

//DECLARE VARIABLE with IDENTIFIER "z" with TYPE "string"
//and ASSIGN a VALUE

var z = "Shaken, not Stirred!"
var a string = ` Bond said:
"Shaken,

not stirred!!!
"`

//STATIC PROGRAMMING LANGUAGE
//VARIABLE DECLARED to hold a VALUE of a certain TYPE

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
