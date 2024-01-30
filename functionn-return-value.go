package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello, Land"
	} else {
		return "Hello" + name
	}
}

func main() {
	result := getHello("Muhammad")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
