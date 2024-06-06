package main

import "fmt"

func main() {
	// length 10, capacity 10
	slice := make([]int, 10)
	fmt.Println(slice)

	// length 0, capacity 10
	slice1 := make([]int, 0, 10)
	fmt.Println(cap(slice1), len(slice1))
}
