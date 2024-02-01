package main

import "fmt"

type Blacklist func(string) bool

func registeredUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked!", name)
	} else {
		fmt.Println("welcome", name)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registeredUser("alif", blacklist)
	registeredUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
