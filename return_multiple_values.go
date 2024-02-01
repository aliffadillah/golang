package main

import "fmt"

func getFullName() (string, string) {
	return "Muhammad", "Fadillah"
}
func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
