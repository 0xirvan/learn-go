package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// Set the number of logical processors to 2
	runtime.GOMAXPROCS(2)

	// buffered channel adalah channel yang memiliki kapasitas
	// channel ini bisa menampung beberapa data sebelum diambil oleh goroutine
	// buffered channel bisa digunakan untuk mengurangi blocking pada goroutine

	messages := make(chan string, 4)

	go func() {
		for {
			i := <-messages
			fmt.Println("Received message", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Sending message", i)
		messages <- fmt.Sprintf("message %d", i)
	}

	time.Sleep(1 * time.Second)
}
