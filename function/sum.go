package main

import "fmt"

func main() {
	// a := 30
	// b := 20
	var a, b int
	fmt.Print("Enter your two numbers:")
	fmt.Scan(&a, &b)

	summation(a, b)
}

func summation(num1 int, num2 int) {
	sum := num1 + num2
	println("The Sum of number is:", sum)
}
