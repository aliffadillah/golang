package main

import (
	"fmt"
	"github.com/aliffadillah/golang/helper"
)

func main() {
	helper.SayHello("Alif")
	fmt.Println("Alif")

	fmt.Println(helper.Application) //bisa diakses diluar package karena menggunakan awalan kapital
	//fmt.Println(helper.version) // tidak bisa karena sebaliknya
}
