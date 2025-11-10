package main

import "fmt"

func changeNumber(num int) { // passed by value, original variable won't change
	num = 5

	fmt.Println("Changed number:", num)
}

// pointer is used to directly change the original variable - passed by reference

func changeOriginalNumber(numPtr *int) { // *int indicates pointer to an int type
	*numPtr = 10 // dereference operator * to get value at the memory address

	fmt.Println("Changed original number:", *numPtr, "from changeOriginalNumber function")
}

func main() {
	number := 1

	changeNumber(number)

	// changeOriginalNumber()
	fmt.Println("Memory address", &number) // & operator to get memory address of variable
	changeOriginalNumber(&number)

	fmt.Println("After changed:", number, "from main function")
}
