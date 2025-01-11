package main

import "fmt"

func main() {
	type student struct {
		name        string
		height      float64
		age         int
		isGraduated bool
		hobbies     []string
	}

	student1 := student{
		name:        "John Doe",
		height:      1.75,
		age:         20,
		isGraduated: false,
		hobbies:     []string{"reading", "swimming"},
	}

	// Format %b untuk menampilkan nilai dalam bentuk biner
	fmt.Printf("%b\n", student1.age)

	// Format %c untuk menampilkan nilai dalam bentuk karakter unicode
	fmt.Printf("%c\n", 65)

	// Format %d untuk menampilkan nilai dalam bentuk berbasis 10
	fmt.Printf("%d\n", student1.age)

	// Format %e untuk menampilkan nilai dalam bentuk notasi ilmiah
	fmt.Printf("%e\n", student1.height)

	// Format %f untuk menampilkan nilai dalam bentuk desimal
	fmt.Printf("%f\n", student1.height)
	fmt.Printf("%.2f\n", student1.height)

	// Format %o untuk menampilkan nilai dalam bentuk oktal
	fmt.Printf("%o\n", student1.age)

	// Format %p untuk menampilkan alamat pointer
	fmt.Printf("%p\n", &student1.name)

	// Format %q untuk menampilkan nilai dalam bentuk string dengan tanda kutip
	fmt.Printf("%q\n", student1.name)

	// Format %s untuk menampilkan nilai dalam bentuk string
	fmt.Printf("%s\n", student1.name)

	// Format %t untuk menampilkan nilai dalam bentuk boolean
	fmt.Printf("%t\n", student1.isGraduated)

	// Format %T untuk menampilkan tipe data
	fmt.Printf("%T\n", student1.name)

	// Format %v untuk menampilkan nilai dalam bentuk default
	fmt.Printf("%v\n", student1)
}
