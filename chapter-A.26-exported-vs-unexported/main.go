package main

import (
	"belajar_golang_level_akses/library"
	"fmt"
)

func main() {
	library.SayHello()
	student := library.Student{Name: "Jhon Doe", Age: 27}
	student.SayHello()

	fmt.Println(library.Person.Name)
	fmt.Println(library.Person.Age)
}
