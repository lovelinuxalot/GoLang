package main

import "fmt"

func main() {
	// the scope of x variable is limited to the if condition itself
	if x := 42; x == 2 {
		fmt.Println(x)
		fmt.Println("True statement")
	}
	// writing multiple statements in a single line requires a ';' to separate the statements
	// Go fmt automatically adds the new statement to new line
	fmt.Println("THis is a statement")
	fmt.Println("There's another one too..")

}
