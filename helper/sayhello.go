package helper

import "fmt"

var version = "1.0"
var Application = "Belajar"

// private
func sayHello(name string) {
	fmt.Println("Hello", name)
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}
