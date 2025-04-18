package main

import (
	"fmt"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		var data = fmt.Sprintf("Data %d", i)
		wg.Add(1)
		go doPrint(&wg, data)
	}

	wg.Wait()
}
