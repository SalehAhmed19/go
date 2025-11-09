package main

import "fmt"

func main() {
	// for loop only construct in Go
	// while like loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// traditional for loop
	for j:=0; j < 5; j++{
		fmt.Println(j)
	}

	// for range loop
	for k:= range 3{
		fmt.Println(k)
	}

	// infinite loop
	for {
		println(1)
	}
}