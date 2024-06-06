package main

import "fmt"

func main() {
	slice := []int{1, 4: 5, 10: 11, 12}
	fmt.Println(slice)
	fmt.Println(len(slice))

	var slc []int
	fmt.Println(slc)
	fmt.Println(slc == nil)

	// append
	slc = append(slc, -5)
	slc = append(slc, -4, -3, -2, -1, 0)
	fmt.Println(slc)
	slc = append(slc, slice...)
	fmt.Println(slc)
}
