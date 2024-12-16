package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Employed bool
}

func NewPerson(name string, age int, employed bool) Person {
	return Person{name, age, employed}
}

func UpdateNameByValue(person Person, newName string) Person {
	person.Name = newName
	return person
}

func UpdateNameByReference(person *Person, newName string) {
	person.Name = newName
}

func MakePersonEmployed(person *Person) {
	person.Employed = true
}

func main() {
	person := NewPerson("John", 30, false)
	fmt.Println(person)

	person = UpdateNameByValue(person, "Jane")
	fmt.Println(person)

	UpdateNameByReference(&person, "Jack")
	fmt.Println(person)

	MakePersonEmployed(&person)
	fmt.Println(person)
}
