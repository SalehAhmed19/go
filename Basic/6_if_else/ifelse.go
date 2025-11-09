package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("You are an adult.")
	}else if age >= 12 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	role:= "admin"
	hasPermission:= true

	if role == "admin" && hasPermission { // && logical AND, || logical OR
		fmt.Println("Access granted to admin panel.")
	}

	// we can declare variable inside if statement
	if name:="Mahin"; name == "Mahin"{
		fmt.Println("Hello, Mahin!")
	} else if name == "Saleh" {
		fmt.Println("Hello, Saleh!")
	}

	// go does not have ternary operator (condition ? value1 : value2)
}