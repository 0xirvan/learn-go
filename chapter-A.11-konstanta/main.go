package main

import "fmt"

func main() {
	// Konstanta adalah variabel yang nilainya tidak dapat diubah
	const firstName = "Jhon"
	const lastName = "Doe"

	fmt.Println("Halo", firstName, lastName)

	// Deklarasi Multi Konstanta
	const (
		name  = "Jhon Doe"
		hobby = "Coding"
	)

	fmt.Println("Nama saya", name, "saya suka", hobby)
}
