package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1  // pass by value
	address3 := &address1 // pass by references

	address2.City = "Bandung"
	address3.City = "Kediri"
	//address3 = &Address{"Subang", "Jawa Barat", "Indonesia"} // not reference, create new data
	*address3 = Address{"Malang", "Jawa Timur", "Indonesia"} // reference, replace first data

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address1)

	alamat1 := new(Address)
	alamat2 := alamat1

	fmt.Println(alamat1)
	alamat2.Country = "Indonesiaa"
	fmt.Println(alamat2)

}
