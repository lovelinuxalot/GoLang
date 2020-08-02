package main

import (
	"fmt"
	"runtime"
)

var x int = 128

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
