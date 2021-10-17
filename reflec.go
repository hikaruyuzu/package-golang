package main

import (
	"fmt"
	"reflect"
)

// struct to make general library

type Sample struct {
	name string `"required":"true" "max"":"10"` // tag
}

type ExampleAgain struct {
	name string `"required"":"true" "max":"10"`
	description string
}

func IsValidate(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Tsukasa"}

	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	example := ExampleAgain{"Kaguya", "oke"}

	fmt.Println(IsValidate(sample))
	fmt.Println(IsValidate(example))

}
