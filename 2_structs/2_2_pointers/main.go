package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	john := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "email@gmail.com",
			zipCode: 94000,
		},
	}

	johnPointer := &john
	johnPointer.updateName("Jimmy")

	john.print()

}

// pointer is being converted into the value (which is the struct) and updated
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
