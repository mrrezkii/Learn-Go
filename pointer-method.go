package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rezki := Man{"Rezki"}
	rezki.Married()
	fmt.Println(rezki.Name)
}
