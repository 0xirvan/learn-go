package main

import "fmt"

func main() {
	// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru

	// deklarasi struct
	type Person struct {
		name string
		age  int
	}

	// deklarasi variabel struct
	var p1 Person // default value

	// set value
	p1.name = "Jhon doe"
	p1.age = 27

	// deklarasi variabel struct + set value
	p2 := Person{name: "Jane Doe", age: 25}

	fmt.Println(p1)
	fmt.Println(p2)

	// Anonymous struct
	// deklarasi anonymous struct
	p3 := struct {
		name string
		age  int
	}{name: "Jhon Doe", age: 27}

	fmt.Println(p3)

	// Slice anonymous struct
	var allStudents = []struct {
		Person
		grade int
	}{
		{Person: Person{"wick", 21}, grade: 2},
		{Person: Person{"ethan", 22}, grade: 3},
		{Person: Person{"bond", 21}, grade: 3},
	}
	for _, student := range allStudents {
		fmt.Println(student)
	}
}
