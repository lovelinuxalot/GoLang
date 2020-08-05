package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	mytruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	mysedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(mytruck)
	fmt.Println(mytruck.doors)
	fmt.Println(mytruck.color)
	fmt.Println(mytruck.fourWheel)
	fmt.Println(mysedan)
	fmt.Println(mysedan.doors)
	fmt.Println(mysedan.color)
	fmt.Println(mysedan.luxury)
}
