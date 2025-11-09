package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch case
	i := 5

	switch i {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	case 4:
		println("Four")
	case 5:
		println("Five")
	default:
		println("Unknown Number")
	}

	// switch case without expression
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default: fmt.Println("It's a workday")
	}

	// type switch
	whoAmI := func (i interface{})  {
		switch t:= i.(type){
			case int:
				fmt.Printf("I'm an integer and my value is %d\n", t) // %d for integer
			case string:
				fmt.Printf("I'm a string and my value is %s\n", t) // %s for string
			default:
				fmt.Printf("I'm of a different type %T\n", t) // %T for type
		}
	}

	whoAmI(10)
}