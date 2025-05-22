package main

import "fmt"

func main() {
	a := 50
	b := 20

	p, q := getNumbers(a, b)
	fmt.Println("Sum is:", p)
	fmt.Println("Multiplication is:", q)

}

func getNumbers(num1 int, num2 int) (int, int) {

	sum := num1 + num2
	mul := num1 * num2

	return sum, mul

}
