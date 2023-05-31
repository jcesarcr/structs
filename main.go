package main

import "fmt"

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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@party.com",
			zipCode: 12345,
		},
	}

	fmt.Println(jim)
	jim.print()
	fmt.Println("\n=========")

	//jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()

	fmt.Println("\n=========")
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
