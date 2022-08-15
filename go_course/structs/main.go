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

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *person) updateEmail(newEmail string) {
	(*p).contact.email = newEmail
}

func main() {
	alex := person{"Alex", "Anderson", contactInfo{"alex@gmail.com", 12345}}
	john := person{firstName: "John", lastName: "Doe", contact: contactInfo{email: "john@gmail.com", zipCode: 12345}}
	fmt.Println(alex, john)

	var asep person
	fmt.Println(asep)

	fmt.Printf("%+v\n", alex)

	alex.firstName = "Alex"
	alex.lastName = "Henderson"

	fmt.Println(alex)

	alexPointer := &alex

	alexPointer.updateName("Putin")

	// you can directly use the value, Go is smart enough to figure out what is required by the function
	// if it requires a value then it will give the value
	// if it requires a pointer then it will give the pointer instead
	alex.updateName("Biden")
	alex.updateEmail("haha@gmail.com")

	alex.print()

}
