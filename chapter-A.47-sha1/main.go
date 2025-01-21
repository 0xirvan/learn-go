package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	// Hash adalah sebuah fungsi yang digunakan untuk menghasilkan nilai hash dari sebuah data
	// Hash digunakan untuk mengamankan data, karena nilai hash tidak bisa diubah

	// Hashing menggunakan SHA1
	text := "This is a secret message"
	sha := sha1.New()
	sha.Write([]byte(text))
	hash := sha.Sum(nil)

	// Hashing menggunakan SHA1 dengan cara yang lebih singkat
	hash2 := sha1.Sum([]byte(text))

	fmt.Println(hash)
	fmt.Println(hash2)

	// Salt
	// Salt adalah data acak yang digunakan sebagai input tambahan untuk fungsi hash
	// Salt digunakan untuk menghasilkan nilai hash yang berbeda meskipun inputnya sama

	// Hashing menggunakan SHA1 dengan salt
	salt := fmt.Sprint(time.Now().UnixNano())
	sha2 := sha1.Sum([]byte(text + salt))
	fmt.Println(sha2)
}
