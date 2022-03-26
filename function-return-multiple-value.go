package main

import "fmt"

func getFullName() (string, string, string) {
	return "Muhammad Rezki", "", "Ananda"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
