package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile("i([a-z])r")

	fmt.Println(regex.MatchString("rezki"))
	fmt.Println(regex.MatchString("ikzer"))
	fmt.Println(regex.MatchString("reZki"))

	fmt.Println(regex.FindAllString("rezki ikzer reZki r3zk1", 10))
}
