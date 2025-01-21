package main

import (
	"fmt"
	"strings"
)

func main() {
	// Fungsi strings contains
	var isExist = strings.Contains("Hello World", "World")
	fmt.Println(isExist) // true

	// Fungsi strings has prefix (digunakan untuk mengecek awalan string)
	var isPrefix = strings.HasPrefix("Hello World", "Hello")
	fmt.Println(isPrefix) // true

	// Fungsi strings has suffix (digunakan untuk mengecek akhiran string)
	var hasSuffix = strings.HasSuffix("Hello World", "World")
	fmt.Println(hasSuffix) // true

	// Fungsi strings count (digunakan untuk menghitung jumlah string yang sama)
	var count = strings.Count("Hello World", "o")
	fmt.Println(count) // 2

	// Fungsi strings index (digunakan untuk mencari posisi string)
	var index = strings.LastIndex("Hello World", "o")
	fmt.Println(index) // 4

	// Fungsi strings replace (digunakan untuk mengganti string)
	var replace = strings.Replace("Hello World", "World", "Golang", 1)
	fmt.Println(replace) // Hello Golang

	// Fungsi strings repeat (digunakan untuk mengulang string)
	var repeat = strings.Repeat("Hello", 3)
	fmt.Println(repeat) // HelloHelloHello

	// Fungsi strings split (digunakan untuk memecah string)
	var split = strings.Split("Hello-World", "-")
	fmt.Println(split)

	// Fungsi strings join (digunakan untuk menggabungkan string)
	var join = strings.Join([]string{"Hello", "World"}, " ")
	fmt.Println(join) // Hello World

	// Fungsi strings to lower (digunakan untuk mengubah string menjadi huruf kecil)
	var toLower = strings.ToLower("Hello World")
	fmt.Println(toLower)

	// Fungsi strings to upper (digunakan untuk mengubah string menjadi huruf besar)
	var toUpper = strings.ToUpper("Hello World")
	fmt.Println(toUpper)
}
