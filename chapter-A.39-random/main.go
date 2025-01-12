package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomStr(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("1", randomizer.Int())
	fmt.Println("2", randomizer.Int())
	fmt.Println("3", randomizer.Int())

	fmt.Println(randomStr(10))
}
