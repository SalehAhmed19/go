package main

import "fmt"

func main() {
	// number sequence of specific length - array
	var arr [4]int
	fmt.Println(len(arr)) // length of array

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40

	// if any element is not assigned, it takes zero value of that type

	fmt.Println(arr) // print entire array

	numbers:= [5]int{1,2,3,4,5} // array literal - single line
	fmt.Println(numbers)

	// 2D array
	nums:=[2][2]int{
		{1,2},
		{3,4},
	}
	fmt.Println(nums)

	/*
	array properties:
	- fixed size, that is predictable
	- Memory optimization
	- Constant time access to elements via index
	*/

	
}