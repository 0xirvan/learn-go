package main

import "fmt"

func main() {
	// Seleksi kondisi menggunakan if, else if, else
	var umur = 20
	if umur > 18 {
		fmt.Println("Umur kamu sudah dewasa")
	} else if umur == 18 {
		fmt.Println("Umur kamu masih muda")
	} else {
		fmt.Println("Umur kamu masih anak-anak")
	}

	// Variabel temporary pada if - else
	var point float32 = 8840.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s sangat baik!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s cukup!\n", percent, "%")
	}

	// Switch case
	var nilai = 69
	switch {
	case nilai >= 90:
		fmt.Println("Nilai kamu A")
	case nilai >= 80:
		fmt.Println("Nilai kamu B")
	case nilai >= 70:
		fmt.Println("Nilai kamu C")
	default:
		{
			fmt.Println("Nilai kamu E")
			fmt.Println("Belajar lagi ya")
		}

	}

	// Keyword fallthrough
	var nilai2 = 70
	switch {
	case nilai2 >= 90:
		fmt.Println("Nilai kamu A")
	case nilai2 >= 80:
		fmt.Println("Nilai kamu B")
	case nilai2 >= 70:
		fmt.Println("Nilai kamu C")
		fallthrough
	// fallthrough digunakan untuk menjalankan case selanjutnya meskipun kondisi tidak terpenuhi
	case nilai2 < 70:
		fmt.Println("Nilai kamu E")
		fmt.Println("Belajar lagi ya")
	}
}
