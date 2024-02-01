package main

import "fmt"

type hasName interface {
	GetName() string
}

func SayHello(value hasName) {
	fmt.Println("hello", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Name
}
func main() {
	person := Person{Name: "Alif"}
	SayHello(person)
	animal := Animal{Name: "Nyonyo"}
	SayHello(animal)
}
