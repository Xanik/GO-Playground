package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{
		email:   "alex@test.com",
		zipCode: 42000,
	}}
	alex.updateName("jimmy")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
