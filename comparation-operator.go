package main

import "fmt"

func main() {
	var name1 = "Rezki"
	var name2 = "Rezki"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 299

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
