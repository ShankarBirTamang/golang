package main

import "fmt"

func arrays() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("thisis the last number", numbers[len(numbers)-1])

	//array of strings
	names := [3]string{"John", "Jane", "Jim"}
	fmt.Println(names)
	fmt.Println("this is the last name", names[len(names)-1])

	//slice
	slice := numbers[1:3]
	fmt.Printf("Slice: %v\n", slice)

	//slice of strings
	sliceNames := names[1:2]
	fmt.Printf("Slice of strings: %v\n", sliceNames)

	//multidimensional array
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("Multidimensional array: %v\n", matrix)
}
