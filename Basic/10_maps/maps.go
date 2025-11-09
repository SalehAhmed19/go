package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps -> hash, object, dictionary
	m := make(map[string]string) // create map with key and value of type string
	// if elements not known at compile time, use make

	// adding key-value pairs
	m["name"] = "Golang"

	// get value by key
	fmt.Println(m["name"])

	// IMP: if key value does not exist, it returns zeroed value of that type
	fmt.Println(m["nonexistent"]) // returns ""

	fmt.Println(len(m)) // length of map

	// deleting key-value pair
	delete(m, "name")

	fmt.Println(m)

	clear(m) // clear the map
	fmt.Println(m)

	// map literal - when elements are known at compile time
	n:= map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(n)

	// check if key exists and get value
	o:=map[string]int{"price": 30, "phone":3} 

	_v, ok:=o["price"] // ok = 30 

	if ok {
		fmt.Println("All okay!")
	} else{
		fmt.Println("Not Okay")
	}

	fmt.Println(_v)

	map1:=map[int]string{1:"a", 2:"b", 3:"c"}
	map2:=map[int]string{1:"a", 2:"b", 3:"c"}

	fmt.Println(maps.Equal(map1, map2))
}