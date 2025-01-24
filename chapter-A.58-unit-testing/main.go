package main

type Kubus struct {
	Sisi float64
}

func (k Kubus) Volume() float64 {
	return k.Sisi * k.Sisi * k.Sisi
}

func (k Kubus) Luas() float64 {
	return k.Sisi * k.Sisi * 6
}

func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}

func main() {

}
