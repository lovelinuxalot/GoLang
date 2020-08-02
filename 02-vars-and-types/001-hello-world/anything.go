package main  		//The first entry-point

import "fmt"

func main() {		// main function. Entrypoint and exit of a program
	fmt.Println("Hello World")
	foo()
	fmt.Println("Something more")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}
func foo() {
	fmt.Println("Im in foo")
}

func bar()  {
	fmt.Println("and then I exited")
}

// Control Flow
// (1) Sequence
// (2) loop; iterative
// (3) conditional