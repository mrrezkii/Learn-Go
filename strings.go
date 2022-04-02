package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Rezki Ananda", "Rezki"))
	fmt.Println(strings.Split("Muhammad Rezki Ananda", " "))
	fmt.Println(strings.ToLower("Muhammad Rezki Ananda"))
	fmt.Println(strings.ToUpper("Muhammad Rezki Ananda"))
	fmt.Println(strings.ToTitle("Muhammad Rezki Ananda"))
	fmt.Println(strings.Trim("     Muhammad Rezki Ananda      ", " "))
	fmt.Println(strings.ReplaceAll("Rezki Rezki Rezki", "Rezki", "Nando"))
}
