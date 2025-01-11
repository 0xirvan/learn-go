package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`Received data"`, data, `"`, "\n")
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout! No activities in 5 seconds")
			break loop
		}
	}
}

func main() {
	messages := make(chan int)

	go sendData(messages)
	retrieveData(messages)
}
