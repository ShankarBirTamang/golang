package go_basics

import "fmt"

func Closure() {
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	//closure with parameters
	doubler := multiplier(2)
	tripler := multiplier(3)
	fmt.Println(doubler(5))
	fmt.Println(tripler(5))

	//closure modifying variables outside its scope
	ModifyVariable()
}

// simple closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// closure with parameters
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

// closure modifying variables outside its scope
// anonymous function can access and modify variables declared in the outer scope
func ModifyVariable() {
	x := 10
	increment := func() {
		x++
	}
	increment()
	increment()
	fmt.Println(x)
}
