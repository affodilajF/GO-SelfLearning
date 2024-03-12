package main

import "fmt"

//import "fmt"
//
//func main() {
//	fmt.Println("Hai")
//	sayHello(Person{"Tasya", "Adinda"})
//}

// ----- interface concept (its implementation processes) in go is a bit different compared to another lang (esp java) ------- //

// HasName this is the interface
// ALL THE METHOD CONTRACT SHOULD BE IMPLEMENTED
type HasName interface {
	getFirstName() string
	getLastName() string
}

// Person this is the (un-manually) implementation of interface
// yes, using a struct
type Person struct {
	firstName string
	lastName  string
}

// the implementation
// once the struct implementing a particular method registered in interface, the whole method should also be implemented
// the struct will be considered implementing the interface
func (person Person) getFirstName() string {
	return person.firstName
}

func (person Person) getLastName() string {
	return person.firstName
}

// this is the method using interface
// the struct passed as "data" should be implementing HasName interface
func sayHello(data HasName) {
	fmt.Println(data.getFirstName())
}
