package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var cust1 Customer
	cust1.Name = "Muhammad Rezki Ananda"
	cust1.Address = "Bandung"
	cust1.Age = 21

	fmt.Println(cust1.Name)
	fmt.Println(cust1.Address)
	fmt.Println(cust1.Age)

	cust1.sayHi("Rezki")

	nando := Customer{
		Name:    "Nando",
		Address: "Kediri",
		Age:     21,
	}

	muhammad := Customer{"Muhammad", "Mekkah", 60}
	fmt.Println(muhammad)

	fmt.Println(nando)
}
