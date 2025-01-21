package main

import (
	"fmt"
	r "regexp"
)

func main() {
	// Regexp adalah package yang digunakan untuk mencari pola dalam sebuah string

	// Contoh penggunaan regexp
	var text = "Hello, my name is John Doe. I am 27 years old. I live in New York."
	var regexp, err = r.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Fungsi matchString digunakan untuk mencari pola dalam sebuah string
	var res1 = regexp.MatchString(text)
	fmt.Println(res1) // true

	// Fungsi findString digunakan untuk mencari pola dalam sebuah string
	var res2 = regexp.FindString(text)
	fmt.Println(res2) // ello

	// Fungsi findAllString digunakan untuk mencari semua pola dalam sebuah string
	var res3 = regexp.FindAllString(text, -1)
	fmt.Println(res3) // [ello my name is ohn oe am years old live in ew ork]

	// Fungsi replaceAllString digunakan untuk mengganti semua pola dalam sebuah string
	var res4 = regexp.ReplaceAllString(text, "what")
	fmt.Println(res4)

	// Fungsi replaceAllStringFunc digunakan untuk mengganti semua pola dalam sebuah string dengan fungsi
	var text2 = "banana burger soup"
	var str = regexp.ReplaceAllStringFunc(text2, func(s string) string {
		if s == "burger" {
			return "pizza"
		}
		return s
	})
	fmt.Println(str) // banana pizza soup

	// Fungsi split digunakan untuk memisahkan string berdasarkan pola
	var regexp2 = r.MustCompile(`[a-b]+`)
	var res5 = regexp2.Split(text2, -1)
	fmt.Printf("%s \n", res5) // [  n n  urger soup]

}
