package main

import "fmt"

func main() {
	name := "Muhammad "

	switch name {
	case "Muhammad":
		fmt.Println("Hello Muhammad")
		fmt.Println("Hello Muhammad")
	case "Alif":
		fmt.Println("Hello Alif")
		fmt.Println("Hello Alif")
	default:
		fmt.Println("Hello GoLand")
		fmt.Println("Hello GoLand")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
