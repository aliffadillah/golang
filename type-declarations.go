package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKTPAlif = "147326849255984"
	var marriedStatus Married = true
	fmt.Println(noKTPAlif)
	fmt.Println(marriedStatus)
}
