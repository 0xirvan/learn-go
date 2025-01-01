package main

import (
	"fmt"
)

func main() {
	// closure
	counter := 0

	increment := func() {
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}
