package main

import "fmt"

func main() {
	name := "Rezki"

	switch name {
	case "Rezki":
		fmt.Println("Hallo Rezki")
	case "Ananda":
		fmt.Println("Hallo Ananda")
	default:
		fmt.Println("Hehehe")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama pas")
	//}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah pas")
	}
}
