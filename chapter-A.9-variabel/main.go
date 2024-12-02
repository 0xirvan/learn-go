package main

import "fmt"

func main() {
	// Manifest typing (Ketika Deklarasi variabel, harus menyertakan tipe data)
	var firstName string = "Jhon"

	var lastName string = "Doe"

	// Type inference (Ketika Deklarasi variabel, tidak perlu menyertakan tipe data)

	var age = 40

	fmt.Printf("Halo nama saya %s %s!,usia saya %d\n", firstName, lastName, age)

	// Deklarasi Multi Variabel
	var name, hobby, isMarried = "Jhon doe", "Coding", true

	fmt.Printf("Nama saya %s, saya suka %s, dan saya sudah menikah: %t\n", name, hobby, isMarried)

	// Deklarasi Variabel tanpa menggunakan var
	variable := "Hello, World!"
	fmt.Println(variable)

	// Variable underscore (_) digunakan untuk menampung nilai yang tidak digunakan
	_ = "Hello, World!"

	// Deklarasi Variabel menggunakan keyword new, yang mengembalikan pointer
	pointer := new(string)
	*pointer = "Ini nilai pointer"
	fmt.Println("Address :", pointer)
	fmt.Println("Value :", *pointer)
}
