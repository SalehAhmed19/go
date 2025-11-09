package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int { // same type for consecutive parameters
	return a - b
}
// function can return multiple values in go

func getLanguages() (string, string, string) {
	return "Go", "Python", "JavaScript"
}

func processIt (f func(a int)int){
	f(1)
}

func process() func(a int) int { // returning a function
	return func(a int) int {
		return 4
	}
}

func main() {
	addition:= add(5, 7)
	subtraction:= sub(3, 4)
	fmt.Println(addition)
	fmt.Println(subtraction)

	// get all languages
	fmt.Println(getLanguages())
	lang1, lang2, _:= getLanguages() // ignore the 3rd return value using _

	fmt.Println(lang1)
	fmt.Println(lang2)
	
	// function as first class citizen - assign function to a variable and pass as argument and return from another function
	fn:= func (a int) int {
		return a * a
	}
	processIt(fn)

	resultFn:= process()
	fmt.Println(resultFn(5))

}