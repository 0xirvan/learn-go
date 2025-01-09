package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = "Rahasia"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	// Tipe any merupakan tipe data yang bisa menampung data apapun
	// Tipe data ini mirip dengan tipe data interface{}
	// Perbedaannya, tipe data any tidak bisa digunakan untuk mengecek tipe data

	data := map[string]any{
		"name":      "Jhon Doe",
		"age":       27,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	// casting data["breakfast"] ke string
	breakfast := strings.Join(data["breakfast"].([]string), ",")
	fmt.Println(breakfast)
}
