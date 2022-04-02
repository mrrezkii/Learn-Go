package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
