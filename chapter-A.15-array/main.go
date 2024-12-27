package main

import "fmt"

func main() {
	// Array
	var names [5]string
	names[0] = "Jhon"
	names[1] = "Doe"
	names[2] = "Foo"
	names[3] = "Bar"
	names[4] = "Leon"

	fmt.Println(names[0], names[1], names[2], names[3], names[4])

	// Inisialisai awal array
	var fruits = [3]string{"apple", "grape", "banana"}

	fmt.Println("Jumlah element = \t", len(fruits))
	fmt.Println(fruits)

	// Insialisasi array tanpa jumlah element
	var numbers = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Jumlah element = \t", len(numbers))
	fmt.Println(numbers)

	// Array multidimensi
	var array2D = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(array2D)

	// Loop array menggunakan for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Element %d : %s\n", i, fruits[i])
	}

	// Loop array menggunakan for range
	for i, fruit := range fruits {
		fmt.Printf("Element %d : %s\n", i, fruit)
	}

}
