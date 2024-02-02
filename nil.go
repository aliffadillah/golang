package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}

}

func main() {
	data := newMap("ALif")
	if data == nil {
		fmt.Println("data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
