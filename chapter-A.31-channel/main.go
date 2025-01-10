package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)
	// Channel adalah tempat komunikasi antar goroutine
	// Channel bisa digunakan untuk mengirim dan menerima data
	// Untuk membuat channel, kita bisa menggunakan fungsi make dan menambahkan tipe data yang akan dikirimkan melalui channel tersebut

	messages := make(chan string)

	for _, each := range []string{"Jhon", "Doe", "Golang"} {
		go func(who string) {
			data := fmt.Sprintf("Hello, %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
