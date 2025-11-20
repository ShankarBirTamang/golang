package main

import "fmt"

func function() {
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)
	sum, product := calculateSumAndProduct(10, 20)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func calculateSumAndProduct(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

// 1. Anonymous function
func outerFunction() {
	innerFunction := func() {
		fmt.Println("This is an anonymous function and it is called inside the outer function")
	}
	innerFunction()
	fmt.Println("This is the outer function")
}

// 2 inline anonymous function
func anonymousFunction() {
	func() {
		fmt.Println("This is an inline anonymous function and it is called inside the anonymous function")
	}()

	fmt.Println("This is the outer function")
}

// 3 Anonymous function with parameters and return value
func anonymousFuncWithParameters() {
	add := func(a, b int) int {
		return a + b
	}
	result := add(10, 20)
	fmt.Println("Result of the addition is", result)
}

// 4. Anonymous callback function
func process(myFn func(int) int) {
	result := myFn(10)
	fmt.Println("Result of the callback function is", result)
}

func callbackFunction() {
	process(func(x int) int {
		return x * 2
	})
	fmt.Println("This is the callback function")
}
