package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)

	duration := time.Since(start)

	fmt.Printf("Duration: %d\n", duration.Milliseconds())
}
