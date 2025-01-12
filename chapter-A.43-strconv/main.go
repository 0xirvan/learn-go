package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi antar tipe data
	/// strvconv adalah package yang digunakan untuk konversi antar tipe data

	// Atoi
	/// Atoi adalah fungsi yang digunakan untuk mengkonversi string ke integer
	var str = "123"

	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("%T : %d\n", num, num)
	}

	// Itoa
	/// Itoa adalah fungsi yang digunakan untuk mengkonversi integer ke string
	str = strconv.Itoa(num)
	if err == nil {
		fmt.Printf("%T : %q\n", str, str)
	}

	// ParseInt
	/// ParseInt adalah fungsi yang digunakan untuk mengkonversi string ke integer dengan bit size tertentu
	str = "123"
	numP, err := strconv.ParseInt(str, 10, 0)
	if err == nil {
		fmt.Printf("%T : %d\n", numP, numP)
	}

	// Formatint
	// FormatInt adalah fungsi yang digunakan untuk mengkonversi integer ke string dengan bit size tertentu

	var numF = int64(24)
	strF := strconv.FormatInt(numF, 2)
	fmt.Printf("%T : %q\n", strF, strF)

	//casting
	/// casting adalah konversi tipe data secara manual
	var numC = 24
	var numC64 = int64(numC)
	fmt.Printf("%T : %d\n", numC64, numC64)

	//type assertion
	/// type assertion adalah konversi tipe data yang dilakukan pada interface

	var data = map[string]any{
		"name":    "wick",
		"age":     27,
		"height":  182.2,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	for _, val := range data {
		switch val := val.(type) {
		case string:
			fmt.Println(val)
		case int:
			fmt.Println(val)
		case float64:
			fmt.Println(val)
		case bool:
			fmt.Println(val)
		case []string:
			fmt.Println(val)
		default:
			fmt.Println(val)
		}
	}

}
