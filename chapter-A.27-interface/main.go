package main

import (
	"fmt"
	"math"
)

type calculate interface {
	area() float64
	circumference() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) getRadius() float64 {
	return c.radius
}

type rectangle struct {
	length float64
}

func (r rectangle) area() float64 {
	return math.Pow(r.length, 2)
}

func (r rectangle) circumference() float64 {
	return r.length * 4
}

// /
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}
func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunDatar calculate

	bangunDatar = rectangle{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas :", bangunDatar.area())
	fmt.Println("keliling :", bangunDatar.circumference())

	bangunDatar = circle{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas :", bangunDatar.area())
	fmt.Println("keliling :", bangunDatar.circumference())
	fmt.Println("jari-jari :", bangunDatar.(circle).getRadius())

	var bangunRuang hitung = &kubus{4}
	luas := bangunRuang.luas()
	keliling := bangunRuang.keliling()
	volume := bangunRuang.volume()
	fmt.Println("===== kubus")
	fmt.Println("luas :", luas)
	fmt.Println("keliling :", keliling)
	fmt.Println("volume :", volume)
	fmt.Printf("alamat memori: %p\n", &bangunRuang)
	fmt.Printf("alamat memori luas: %p\n", &luas)
	fmt.Printf("alamat memori keliling: %p\n", &keliling)
	fmt.Printf("alamat memori volume: %p\n", &volume)

}
