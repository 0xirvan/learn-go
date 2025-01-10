package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMessage(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	// Goroutine adalah sebuah cara untuk membuat sebuah proses yang berjalan secara asynchronous
	// Goroutine ini akan berjalan secara bersamaan dengan proses utama
	// Untuk membuat goroutine, kita bisa menggunakan kata kunci go diikuti dengan nama fungsi yang ingin dijalankan sebagai goroutine
	go printMessage(5, "Hello")
	go printMessage(5, "World")

	time.Sleep(time.Second)
}
