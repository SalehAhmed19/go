package main

import "fmt"

func main() {
	var name string = "Golang" // variable declaration with type - explicit
	// var age = 10 - variable declaration without type - implicit
	var age int = 10
	var isAdult bool = true
	
	// shorthand syntax for variable declaration and assignment
	fullName:= "Go Developer"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isAdult)
	fmt.Println(fullName)
}