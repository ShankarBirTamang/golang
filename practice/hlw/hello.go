package hello_world

import "fmt"

func Hello() {
	country, capital := "Nepal", "Kathmandu"

	var (
		age       int     = 20
		isStudent bool    = true
		height    float64 = 1.75
	)
	message, message2 := message(country, capital, age, isStudent, height)
	fmt.Println(message)
	fmt.Println(message2)
}

func message(country, capital string, age int, isStudent bool, height float64) (string, string) {
	return fmt.Sprintf("Hello, %s! Capital is %s\n", country, capital), fmt.Sprintf("You are %d years old, %t years old and %f meters tall!\n", age, isStudent, height)
}
