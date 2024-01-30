package main

import "fmt"

func getFullName() (string, string, string) {
	return "Muhammad", "Alif", "Fadillah"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

// "_ (underscore)" untuk ignore suatu variable
