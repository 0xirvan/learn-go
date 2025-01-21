package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello World"
	url := "https://irvan.xyz"

	// Encode
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded:", string(decoded))

	// Penerapan fungsi encode dan decode
	encoded2 := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded2, []byte(data))
	fmt.Println("Encoded 2:", string(encoded2))

	decoded2 := make([]byte, base64.StdEncoding.DecodedLen(len(encoded2)))
	_, err := base64.StdEncoding.Decode(decoded2, encoded2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Decoded 2:", string(decoded2))

	// URL Encoding
	encodedURL := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Encoded URL:", encodedURL)

	decodedURL, _ := base64.URLEncoding.DecodeString(encodedURL)
	fmt.Println("Decoded URL:", string(decodedURL))
}
