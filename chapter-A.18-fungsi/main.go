package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Random number between 1 and 100:", randomWithRange(1, 100))
	fmt.Println("Random number between 10 and 50:", randomWithRange(10, 50))
	fmt.Println("Random number between 100 and 200:", randomWithRange(100, 200))
}

func randomWithRange(min, max int) int {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	return randomizer.Int()%(max-min) + min
}
