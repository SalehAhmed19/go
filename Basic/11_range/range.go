package main

import "fmt"

func main() {
	// range -> loop over array, slice, string, map, channel
	nums := []int{1, 2, 3, 4, 5}

	// basic for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	sum:=0
	// for range loop
	for idx, num:= range nums{
		sum = sum + num
		fmt.Println("Index:", idx, "Value:", num)
	}
	fmt.Println(sum)

	m:=map[string]int{"a":1, "b":2, "c":3}

	for k, v:= range m{
		fmt.Println(k, v)
	}

	// i = starting byte of rune
	for i, c:= range "Golang"{
		fmt.Println(i, string(c)) // without string() it will print the unicode point rune
	}
}