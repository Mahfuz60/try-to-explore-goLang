package mathlib

import "fmt"

func Addition(a int, b int) {
	c := a + b
	fmt.Println("Addition of ", a, "and", b, "is:", c)
}

func Subtraction(a int, b int) {
	d := a - b

	fmt.Println("Subtraction of", a, "and", b, "is:", d)
}

func Multiplication(a int, b int) {
	e := a * b
	fmt.Println("Multiplication of", a, "and", b, "is:", e)
}
func Division(a int, b int) {
	f := a / b
	fmt.Println("Division of", a, "and", b, "is:", f)
}

func Modulus(a int, b int) {
	g := a % b
	fmt.Println("Modulus of", a, "and", b, "is:", g)

}
