package main

import "fmt"

func main() {
	// pointer adalah reference ke alamat memori
	// variabel pointer berarti variabel yang menyimpan alamat memori

	// deklarasi variabel pointer
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	change(&numberA, 10)
	fmt.Println("numberA (value) :", numberA)

}

// Parameter Pointer
func change(original *int, value int) {
	*original = value
}
