package main

import "fmt"

func getGoodBye(name string) string {
	return "GoodBye" + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Rezki")
	fmt.Println(result)
}
