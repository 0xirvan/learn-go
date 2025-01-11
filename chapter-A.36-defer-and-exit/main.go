package main

import (
	"fmt"
	"os"
)

func main() {
	// defer berfungsi untuk menjalankan kode setelah kode sebelumnya selesai dijalankan
	// defer akan dijalankan sebelum kode diakhir fungsi
	func() {
		defer println("world")
		println("hello")
	}()

	// Kombinasi defer dab IIFE

	number := 3

	if number == 3 {

		func() {
			defer fmt.Println("Halo 3")
			os.Exit(1)
			fmt.Println("Halo 1")
		}()

	}

	fmt.Println("Halo 2")
}
