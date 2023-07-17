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
	me := person{
		firstName: "nicholas",
		lastName:  "neuenschwander",
		contactInfo: contactInfo{
			email:   "nicholasneuenschwander@gmail.com",
			zipCode: 30033,
		},
	}

	// mePointer := &me
	// mePointer.updateName("nick")
	me.updateName("nick") // this is used more in practice
	me.print()
	// var me person
	// me.firstName = "nicholas"
	// me.lastName = "neuenschwander"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
