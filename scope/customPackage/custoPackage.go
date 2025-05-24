package main

import (
	"fmt"

	"example.com/mathlib"
)

func main() {

	fmt.Println("Showing custom Package!")
	mathlib.Addition(10, 20)
	mathlib.Subtraction(50, 20)
	mathlib.Multiplication(10, 20)
	mathlib.Division(100, 20)
	mathlib.Modulus(100, 22)

}
