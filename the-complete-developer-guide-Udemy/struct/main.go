package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	fisrtName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		fisrtName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFisrtName string) {
	(*pointerToPerson).fisrtName = newFisrtName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
