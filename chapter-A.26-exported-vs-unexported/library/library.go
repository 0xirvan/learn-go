package library

import "fmt"

// SayHello adalah fungsi publik yang bisa diakses dari luar package
func SayHello() {
	fmt.Println("Hello, my name is Jhon Doe")
	damn()
}

// damn adalah fungsi private yang hanya bisa diakses dari dalam package
func damn() {
	fmt.Println("Damn!")
}

type Student struct {
	Name string
	Age  int
}

func (s Student) SayHello() {
	fmt.Printf("Hello, my name is %s, my age is %d years old\n", s.Name, s.Age)
}

var Person = struct {
	Name string
	Age  int
}{}

func init() {
	Person.Name = "Jhon Doe"
	Person.Age = 27

	fmt.Println("Library initialized")
}
