package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name is", customer.name)
}
func main() {
	var alif Customer
	alif.name = "Muhammad Alif Fadillah, "
	alif.address = "Indonesia, "
	alif.age = 19

	fmt.Println(alif)

	muhammad := Customer{
		name:    "ALif",
		address: "Indonesia",
		age:     19,
	}
	fmt.Println(muhammad)

	fadillah := Customer{"Al", "indoensia", 19}
	fmt.Println(fadillah)

	fadillah.sayHello("112")
}
