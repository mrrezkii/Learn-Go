package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke-", counter)
	//	counter++
	//}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	slice := []string{"Muhammad", "Rezki", "Ananda"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range slice {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Muhammad Rezki Ananda"
	person["address"] = "Bandung"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
