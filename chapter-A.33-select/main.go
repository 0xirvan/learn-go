package main

func getAverage(numbers []int, ch chan float64) {
	var sum int
	for _, each := range numbers {
		sum += each
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max int
	for _, each := range numbers {
		if each > max {
			max = each
		}
	}
	ch <- max
}

func main() {
	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}

	var ch1 = make(chan float64)
	var ch2 = make(chan int)

	go getAverage(numbers, ch1)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			println("Avg \t:", avg)
		case max := <-ch2:
			println("Max \t:", max)
		}
	}
}
