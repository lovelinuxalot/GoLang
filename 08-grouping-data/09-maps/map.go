package main

import "fmt"

func main() {
	m := map[string]string{
		"Jedi Name": "Anakin Skywalker",
		"Sith Name": "Darth Vader",
	}

	fmt.Println(m)
	fmt.Println(m["Sith Name"])

	age_m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(age_m)
	fmt.Println(age_m["James"])
	fmt.Println(age_m["Moneypenny"])
	fmt.Println(age_m["Q"]) // This will return ZERO value, if the key do not exist for int

	v, ok := age_m["Q"]
	fmt.Println(v)
	fmt.Println(ok) // this will return FALSE, if VALUE do not exist

	// Following conditional check can be done to ensure value exists and do stuff

	if v, ok := age_m["Q"]; ok {
		fmt.Println("This prints the value from IF condition", v)
	}

	if v, ok := age_m["Moneypenny"]; ok {
		fmt.Println("This prints the value from IF condition", v)
	}
}
