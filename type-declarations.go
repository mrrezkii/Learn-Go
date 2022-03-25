package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPRezki NoKTP = "1202190044"
	var marriedStatus Married = true
	fmt.Println(noKTPRezki)
	fmt.Println(marriedStatus)
}
