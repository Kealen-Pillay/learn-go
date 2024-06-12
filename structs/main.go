package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	contact ContactInfo	// embedded struct
}

type ContactInfo struct {
	email string
	zip int
}

func main() {
	p := Person{
		firstName: "Kealen",
		lastName: "Pillay",
		contact: ContactInfo{
			email: "test@test.com",
			zip: 1,
		},
	}
	p.setName("Kai")
	fmt.Printf("\n%+v", p)
}

func (p Person) print() {
	fmt.Printf("\n%+v", p)
}

func (p *Person) setName(firstName string) {
	p.firstName = firstName
}