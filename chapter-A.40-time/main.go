package main

import (
	"fmt"
	"time"
)

func main() {
	// Waktu sekarang
	time1 := time.Now()
	fmt.Printf("time 1 %v\n", time1)
	fmt.Printf("year 1 %v\n", time1.Year())
	fmt.Println(time.Now().Year() - 1)

	// waktu yang di tentukan sendiri
	time2 := time.Date(2022, 10, 10, 10, 10, 10, 10, time.Local)
	fmt.Printf("time 2 %v\n", time2)

	//Parsing string ke time
	time3, _ := time.Parse(time.RFC3339, "2022-10-10T10:10:10Z")
	fmt.Printf("time 3 %s\n", time3.Format(time.RFC3339))

	time4 := time.Now().Format(time.Kitchen)
	fmt.Printf("time 4 %v\n", time4)

ajg:
	for {
		if timeNow := time.Now().Format(time.Kitchen); timeNow == "8:40AM" {
			fmt.Println("Waktunya bangun")
			break ajg
		} else {
			fmt.Println("Belum waktunya bangun")
			time.Sleep(1 * time.Second)
		}
	}
}
