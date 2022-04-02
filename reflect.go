package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}

	return true
}

func main() {
	sample := Sample{"Rezki"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(structField.Name)
	fmt.Println(structField.Tag.Get("required"))
	fmt.Println(structField.Tag.Get("max"))

	fmt.Println(IsValid(sample))
	sample.Name = ""
	fmt.Println(IsValid(sample))
}
