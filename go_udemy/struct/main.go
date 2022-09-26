package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// chirag := person{
	// 	firstName: "Chirag",
	// 	lastName:  "Gupta"}
	// fmt.Println(chirag)

	// var axl person
	// fmt.Println(axl)
	// fmt.Printf("%+v\n", chirag)
	// fmt.Printf("%+v\n", axl)
	// axl.firstName = "Axl"
	// axl.lastName = "Heimdallson"
	// fmt.Println(axl)

	chirag := person{
		firstName: "chirag",
		lastName:  "Gupta",
		contact: contactInfo{
			email:   "cg@g.com",
			pincode: 102991,
		},
	}
	// fmt.Printf("%+v\n", chirag)
	chirag.print()
	(&chirag).updateFirstName("Chirag")
	chirag.print()
}

func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
