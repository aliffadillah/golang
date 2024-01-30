package main

import "fmt"

func namaLengkap() (firstName string, middleName string, lastName string) {
	firstName = "Muhammad"
	middleName = "Alif"
	lastName = "Fadillah"
	return
}

func main() {
	a, b, c := namaLengkap()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
