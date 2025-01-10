package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `required:"true"`
}

func isValidData(data interface{}) bool {
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
	p1 := person{"John Doe"}
	number := 10
	reflectValue := reflect.ValueOf(number)

	fmt.Println(reflectValue.Type())
	fmt.Println(reflect.TypeOf(number))

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel adalah", reflectValue.Int())
	}
	fmt.Println(isValidData(p1))
}
