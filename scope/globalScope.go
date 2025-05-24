package main

import "fmt"

//global scope variable
var (
	a = 20
	b = 30
)

func add(x int, y int) {
	z := x + y
	fmt.Println("Sum is:", z)
}

func main() {
	//local scope variable
	var p int = 30
	var q int = 40

	add(a, b)
	add(a, q)
	add(p, a)

}
