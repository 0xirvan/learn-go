package main

import (
	"fmt"
	"strings"
)

func main() {
	// fungsi variadic
	result := sumAll(10, 10, 10, 10, 10)
	println(result)

	// slice sebagai argumen variadic
	numbers := []int{10, 10, 10, 10, 10}
	result = sumAll(numbers...)
	println(result)

	sayHello("Irvan", "Makan", "Ngoding")
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

// fungsi dengan param biasa dan variadic
func sayHello(name string, hobbies ...string) {
	hobbiesAsStr := strings.Join(hobbies, ",")

	fmt.Printf("Hello nama saya %s\nhobi saya %s\n", name, hobbiesAsStr)
}
