package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("Enter the value:")
	fmt.Scanln(&a, &b)

	result1 := sum(a, b)

	fmt.Println("The Sum of Values:", result1)

	result2 := sub(a, b)

	fmt.Println("The Subtraction of Values:", result2)

	result3 := multiply(a, b)
	fmt.Println("The Multiplication of Values:", result3)

	result4 := divide(a, b)

	fmt.Println("The division of Values:", result4)

	result5 := mod(a, b)
	fmt.Println("The Modules of Values:", result5)

}

func sum(num1 int, num2 int) int {

	result1 := num1 + num2

	return result1

}

func sub(num1 int, num2 int) int {

	result2 := num1 - num2

	return result2

}

func multiply(num1 int, num2 int) int {
	result3 := num1 * num2

	return result3
}

func divide(num1 int, num2 int) int {

	result4 := num1 / num2
	return result4
}

func mod(num1 int, num2 int) int {
	result5 := num1 % num2
	return result5
}
