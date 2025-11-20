package main

import "fmt"

func variables() {
	y := 10
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Println(*&y)
	x := &y
	fmt.Println(x)
	fmt.Println(*x)
	var name string = "John"
	fmt.Printf("Name: %s\n", name)

	var district, province string = "Sunsari", "Koshi"
	fmt.Printf("District: %s, Province: %s\n", district, province)

	var (
		age       int     = 20
		isStudent bool    = true
		height    float64 = 1.75
	)
	fmt.Printf("Age: %d, Is Student: %t, Height: %f\n", age, isStudent, height)

	var defaultInt int
	var defaultString string
	var defaultFloat float64
	var defaultBool bool
	fmt.Printf("Default Int: %d, Default String: %s, Default Float: %f, Default Bool: %t\n", defaultInt, defaultString, defaultFloat, defaultBool)
}
