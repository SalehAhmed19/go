package main

import "fmt"

const age int = 10 // constant declaration outside function
func main() {
	const name string = "Golang" // constant declaration with type - explicit
	fmt.Println(age)

	const ( // multiple constant declarations
		isAdult bool = true
		height float64 = 5.9
	)
}