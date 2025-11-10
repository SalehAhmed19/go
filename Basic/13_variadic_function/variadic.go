package main

import "fmt"

func sum(nums ...int) int { // ... indicates variadic parameter - slice of ints
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func greet(users ...interface{}) { // interface{} can accept any type
	for _, u := range users {
		fmt.Println("Hello", u)
	}
}

func main() {
	fmt.Println(1, 2, 3, 4, "hello") // built-in variadic function - n number of arguments

	result := sum(3, 4, 5, 6)

	fmt.Println("Sum is:", result)

	greet(1, "Hi", true)

	numbers := []int{10, 20, 30, 40}
	result2 := sum(numbers...) // unpacking slice to variadic function

	fmt.Println("Sum of slice is:", result2)
}
