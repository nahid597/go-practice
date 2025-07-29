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

	testPerson := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "a@a.com",
			zipCode: 12345,
		},
	}

	testPersonPointer := &testPerson
	testPersonPointer.updateName("Jane updates")
	fmt.Printf("%+v", testPerson)

}

func (pointerPerson *person) updateName(firstName string) {
	(*pointerPerson).firstName = firstName
}
