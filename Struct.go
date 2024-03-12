package main

import "fmt"

//func main() {
//
//	fmt.Println("Hai")
//
//	person1 := PersonA{"Tasya", 12}
//	fmt.Println(person1.name)
//	fmt.Println(person1.age)
//
//	person1.getName()
//	person1.changeName()
//	person1.getName()
//
//	fmt.Println("Hai")
//	fmt.Println(person1)
//
//}

type PersonA struct {
	name string
	age  int
}

// adding method for a struct
func (person PersonA) getName() {
	fmt.Print(person.name)
	fmt.Println()
}

func (person *PersonA) changeName() {
	person.name = person.name + "Addi"
}

// kalo semua fungsinya pointer ya pointer semua, kalo bukan ya egkkk, idk kenapah
