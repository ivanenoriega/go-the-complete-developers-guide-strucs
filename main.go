package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jack")
	jim.print()
}

func (p person) updateName(n string) {
	p.firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
