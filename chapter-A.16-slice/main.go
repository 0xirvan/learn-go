package main

import "fmt"

func main() {
	// Array - Fixed size
	var numbers [3]int = [3]int{1, 2, 3} // Size must be specified

	// Slice - Dynamic size
	var scores []int = []int{1, 2, 3} // No size specified
	scores = append(scores, 4)        // Can grow dynamically
	fmt.Println(scores)
	fmt.Println(cap(scores))
	// Reference behavior demonstration
	array1 := [3]int{1, 2, 3}
	array2 := array1 // Creates a copy
	array2[0] = 10   // Doesn't affect array1

	slice1 := []int{1, 2, 3}

	slice2 := slice1 // Creates a reference
	slice2[0] = 10   // Affects slice1 too

	fmt.Println("Arrays:", array1, array2) // Output: [1 2 3] [10 2 3]
	fmt.Println("Slices:", slice1, slice2) // Output: [10 2 3] [10 2 3]

	// Capacity and length
	fmt.Printf("Array length: %d\n", len(numbers))
	fmt.Printf("Slice length: %d, capacity: %d\n", len(scores), cap(scores))
}
