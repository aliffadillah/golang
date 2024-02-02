package main

import (
	"fmt"
	"github.com/aliffadillah/golang/database"
	_ "github.com/aliffadillah/golang/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
