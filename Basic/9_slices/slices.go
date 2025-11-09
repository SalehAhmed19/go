package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices are dynamic arrays - most used data structure in Go + useful methods
	// uninitialized slice is nil slice - nil means null in go
	var nums []int

	fmt.Println(nums)

	// make function creates slices, maps and channels -> make(type, length, capacity)
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

	// copy function copies elements from source slice to destination slice
	var distSlice = make([]int, 0, 5)
	distSlice = append(distSlice, 1,2,3,4,5)
	var copyDist = make([]int, len(distSlice))


	// copy(destination, source)
	copy(copyDist, distSlice)

	fmt.Println("disSlice: ", distSlice, "copySlice: ", copyDist)

	// slice operators
	numberList:=[]int{1,2,3}

	fmt.Println(numberList[0:2]) // slicing from index 0 to 2 (excluding 2)
	fmt.Println(numberList[:2]) // from start to index 2
	fmt.Println(numberList[1:]) // from index 1 to end
	fmt.Println(numberList[:]) // entire slice

	// slice package
	numsList1:=[]int{5,3,8,1,2}
	numsList2:=[]int{10,20,30,40,50}

	fmt.Println(slices.Equal(numsList1, numsList2)) // check equality

	// 2D slices
	numbs:=[][]int{
		{1,2,3},
		{4,5,6},
	}

	fmt.Println(numbs)
}