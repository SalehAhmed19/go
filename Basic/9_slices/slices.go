package main

import "fmt"

func main() {
	// slices are dynamic arrays - most used data structure in Go + useful methods
	// uninitialized slice is nil slice - nil means null in go
	var nums []int

	fmt.Println(nums)

	var numbers = make([]int, 2) // create slice of int with initial length 5 | third param is capacity (optional)
	numbers[0]=10

	// capacity is the maximum size the slice
	fmt.Println(cap(numbers))

	fmt.Println(numbers)

	numbers = append(numbers, 1) // append element to slice
	numbers = append(numbers, 2, 3, 4, 5) // append multiple elements - capacity doubles when exceeded
	fmt.Println(numbers)

	numberSlice:=[]int{10,20,30,40,50} // slice literal
	fmt.Println(numberSlice)



}