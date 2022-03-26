package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Muhammad Rezki Ananda",
		"address": "Bandung",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address "])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["wrong"] = "haa"

	delete(book, "wrong")

	fmt.Println(book)
}
