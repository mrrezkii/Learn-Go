package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	namedFiltered := filter(name)
	fmt.Println("Hello", namedFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name

}

func main() {
	sayHelloWithFilter("Rezki", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
