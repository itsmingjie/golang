package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	// In order of field definition (not recommended)
	alex := person{"Alex", "Anderson"}

	// OR: Specifc definition
	john := person{firstName: "John", lastName: "Doe"}

	// OR: Zero value assignment
	var joe person

	// will throw error because variables not used
}
