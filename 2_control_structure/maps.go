package control_structure

import "fmt"

func Maps() {
	person := map[string]int{
		"John": 20,
		"Jane": 21,
		"Jim":  22,
	}
	fmt.Println(person)
	fmt.Println(person["John"])

	person["Sankalp"] = 29
	fmt.Printf("Person after addition: %v\n", person)
	delete(person, "Jim")
	fmt.Printf("Person after deletion: %v\n", person)

	for name, age := range person {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	for _, age := range person {
		fmt.Printf("Age: %d\n", age)
	}

	for name := range person {
		fmt.Printf("Name: %s\n", name)
	}

	//check if a key exists in the map
	key := "Sankalp"
	if age, exists := person[key]; exists {
		fmt.Printf("Person with name %s exists in the map and the age is %d\n", key, age)
	} else {
		fmt.Printf("Person with name %s does not exist in the map\n", key)
	}

	//get the length of the map
	fmt.Printf("Length of the map: %d\n", len(person))

	//clear the map
	person = make(map[string]int)
	fmt.Printf("Person after clearing: %v\n", person)
}
