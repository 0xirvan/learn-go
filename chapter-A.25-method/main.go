package main

import "fmt"

// method adalah fungsi yang menempel pada tipe data, MISAL: struct

type Person struct {
	name string
	age  int
}

func (p Person) sayHello() string {
	return fmt.Sprintln("Hello, my name is", p.name)
}

func (p Person) sayAge() string {
	return fmt.Sprintln("I am", p.age, "years old")
}

// method pointer
func (p *Person) saySomething() string {
	p.name = "Master"
	return fmt.Sprintln("Hello, my name is", p.name)
}

func main() {
	var p1 Person
	p1.name = "Jhon Doe"
	p1.age = 27

	fmt.Print(p1.sayHello())
	fmt.Print(p1.sayAge())

	fmt.Print(p1.saySomething())
	fmt.Print(p1.sayHello())
}
