package main

import "fmt"

func main() {
	counter := 0
	name := "Rezki"

	increment := func() {
		name := "Nando"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(name)

	fmt.Println(counter)
}
