package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	alif := Man{"Alif"}
	alif.Married()
	fmt.Println(alif.Name)
}
