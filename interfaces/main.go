package main

import "fmt"

type Person struct{
	Name string
}

type Robot struct{
	Model string
}

type Animal struct{
	Name string
}

type Speaker interface{
	Speak() string
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hi! My name is %s.", p.Name)
}

func (r Robot) Speak() string {
	return fmt.Sprintf("Hello, I am robot model %s.", r.Model)
}

func (a Animal) Speak() string {
	return "this is a test"
}

func Introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	person := Person{Name: "Alice"}
	robot := Robot{Model: "RX-78"}
	animal := Animal{Name: "Cloe"}
	
	speakers := []Speaker{person, robot, animal}
	
	for _, s := range speakers {
		Introduce(s)
	}
}
