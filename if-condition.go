package main

import "fmt"

func main() {
	name := "Muhammad"
 
	if name == "Muhammad" {
		fmt.Println("Hello", name)
	} else if name == "Fadillah" {
		fmt.Println("Hello", name)
	} else if name == "Alif" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hello, Go!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
