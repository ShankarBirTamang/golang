package main

import "fmt"

/*

In Go, you can attach functions to a struct.
When you do that, those functions become methods.

Example struct:
type Person struct {
    Name string
    Age  int
}
Normally we do : pass the *Person as an argument to the function.

func modifyPersonName(p *Person) {
    p.Name = "New Name"
}



*/

//turning the function into a method
func (p *Person) ModifyName() {
	p.Name = "New Name"
}

//Here:
//(p *Person) → method receiver
//- p acts like the variable representing that Person
//- This tells Go: This method belongs to the Person struct

func methodReceiverExample() {
	person := Person{}
	person.Name = "Sankar"
	fmt.Printf("Person name before modification: %s\n", person.Name)
	person.ModifyName()
	fmt.Printf("Person name after modification: %s\n", person.Name)
}
