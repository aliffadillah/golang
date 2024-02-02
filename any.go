package main

import "fmt"

func Ups() interface{} { //goLang versi terbaru menggunakan any
	return 1
}

func main() {
	var kosong = Ups() //tambahkan "any" setelah variable "kosong"
	fmt.Println(kosong)

}
