package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	area, circumference := calculate(diameter)

	fmt.Printf("Luas lingkaran\t\t: %.2f\n", area)
	fmt.Printf("Keliling lingkaran\t: %.2f\n", circumference)
}

func calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)
	circumference := math.Pi * d

	return area, circumference
}

// Predefined return value

func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}
