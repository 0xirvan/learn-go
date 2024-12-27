package main

import "fmt"

func main() {
	// Perulangan for
	for i := 0; i < 5; i++ {
		fmt.Printf("Angka: %d\n", i)
	}

	//Perulangan for tanpa argumen
	var i int = 0
	for {
		fmt.Printf("Angka: %d\n", i)
		i++
		if i == 9 {
			break
		}
	}

	// Perulangan dengan for range
	var arr = [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// Perulangan keyword break and continue
	for i := 1; i <= 20; i++ {
		if i%2 == i {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Printf("Angka %d\n", i)
	}

	// Perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

	// Label
	/* Di perulangan bersarang, break dan continue akan berlaku pada blok
	perulangan di mana ia digunakan saja. */
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}

			fmt.Printf("matriks [ %d ][ %d ] \n", i, j)
		}
	}

}
