// structs are user defined types that allow you to store a collection of values of different types
package control_structure

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street string
	City   string
	State  string
}

func Structs() {
	//nested struct
	person := Person{
		Name:    "John",
		Age:     20,
		Address: Address{Street: "123 Main St", City: "Anytown", State: "CA"},
	}
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Person name: %s\n", person.Name)
	fmt.Printf("Person age: %d\n", person.Age)
	fmt.Printf("Person address: %+v\n", person.Address)
	fmt.Printf("Person address street: %s\n", person.Address.Street)
	fmt.Printf("Person address city: %s\n", person.Address.City)
	fmt.Printf("Person address state: %s\n", person.Address.State)
}

// anonymous struct
func AnonymousStruct() {
	person := struct {
		Name    string
		Age     int
		Address struct {
			Street string
			City   string
			State  string
		}
	}{
		Name:    "John",
		Age:     20,
		Address: Address{Street: "123 Main St", City: "Anytown", State: "CA"},
	}
	fmt.Printf("Anonymous Person: %+v\n", person)
	fmt.Printf("Anonymous Person name: %s\n", person.Name)
	fmt.Printf("Anonymous Person age: %d\n", person.Age)
	fmt.Printf("Anonymous Person address: %+v\n", person.Address)
	fmt.Printf("Anonymous Person address street: %s\n", person.Address.Street)
	fmt.Printf("Anonymous Person address city: %s\n", person.Address.City)
	fmt.Printf("Anonymous Person address state: %s\n", person.Address.State)
}

/*
In Go:
- Structs are passed by value by default

This means:
- When you pass a struct into a function → Go makes a copy
- That copy exists only inside the function
- Changing it does not affect the original struct outside
*/

func PassByValue() {
	updateAge := func(p Person) {
		p.Age = 50 // changes only the copy
	}

	p := Person{Name: "Sankar", Age: 25}
	updateAge(p)

	fmt.Printf("Person age: %d\n", p.Age) // still 25 (NOT changed)
}

// If you want reference behavior → pass a pointer to the struct
func PassByReference() {
	updateAge := func(p *Person) {
		p.Age = 50 // changes the original struct
	}

	p := Person{Name: "Sankar", Age: 25}
	updateAge(&p)

	fmt.Printf("Person age: %d\n", p.Age) // now 50 (changed)
}
