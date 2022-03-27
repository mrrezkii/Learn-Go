package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Rezki"
	lastName = "Ananda"
	return
}

func main() {
	firstName, _, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
