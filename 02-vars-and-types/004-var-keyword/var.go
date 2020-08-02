package main

import "fmt"

//DECLARE and ASSIGN == Initialization
var y = 43

//DECLARE a VARIABLE with IDENTIFIER "z"
//with TYPE "int"
//ASSIGN the ZERO VALUE of TYPE "int" to "z"

//ZERO VALUE is "false" for booleans, 0 for integers,
//0.0 for floats, "" for strings,
//nil for pointers, functions, slices, interfaces, channels and maps

var z int

func main() {

	//Short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println("The VALUE of X: ", x)
	foo()
	y = 45
	fmt.Println("This is VALUE of Y from 'main': ", y)
	fmt.Println("The VALUE of Z: ", z)
}

func foo() {
	fmt.Println("The VALUE of Y from 'foo' function: ", y)
}
