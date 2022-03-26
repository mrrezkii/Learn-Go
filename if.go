package main

import "fmt"

func main() {
	var name = "Ananda"

	if name == "Rezki" {
		fmt.Println("Hello Rezki")
	} else if name == "Ananda" {
		fmt.Println("Hello Ananda ")
	} else {
		fmt.Println("Hi, boleh kenalan ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
