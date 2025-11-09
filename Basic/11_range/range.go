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


}