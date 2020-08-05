package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m)

	m["Q"] = 25
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Q")
	fmt.Println(m)

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println(v)
		delete(m, "Moneypenny")
	}
	fmt.Println(m)
}
