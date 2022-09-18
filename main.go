package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// func Add(){

// } // capital function name is PUBLIC, else PRIVATE

func main() {
	fmt.Println("Hello, World")

	// var assignedButNotUsed = 4 //throws error
	var num1, num3 int
	var num2 = 3
	num1, num3 = 2, 5
	const constant = 4
	// constant = 4
	num4 := 8

	var result = num1 + num2 + num3 + num4
	fmt.Println(result)

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("sum func: ", add(4, 3))

	var name string = "chirag"
	fmt.Println(name)
}
