package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type business struct {
	name string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	business
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Black",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
		business: business{
			name: "JimsShootingSupplies",
		},
	}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
