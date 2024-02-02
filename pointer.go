package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// PASS BY VALUE
	//address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	//address2 := address1
	//
	//address2.City = "Bandung"
	//fmt.Println(address1)
	//fmt.Println(address2)

	//PASS BY REFERENCE
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
}
