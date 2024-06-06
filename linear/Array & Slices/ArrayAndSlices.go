package main

import "fmt"

func main() {
	// Create a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial slice:", numbers)

	// Append elements
	numbers = append(numbers, 6)
	fmt.Println("After appending 6:", numbers)

	// Delete element at index 2
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println("After deleting element at index 2:", numbers)

	// Slicing
	subSlice := numbers[1:4]
	fmt.Println("Sub-slice from index 1 to 4:", subSlice)

	// Capacity
	fmt.Println("Length of sub-slice:", len(subSlice))
	fmt.Println("Capacity of sub-slice:", cap(subSlice))
}
