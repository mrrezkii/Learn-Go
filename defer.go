package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // tetap dipanggil meskipun error, meskipun ada panic function juga
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
