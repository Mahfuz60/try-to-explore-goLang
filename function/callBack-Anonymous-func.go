package main

import "fmt"

//callBack or invoke function
//standard or name function
func add(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	//Anonymous function
	//Immediately Invoked Function Expression(IIFE)

	func(x, y int) {
		z := x + y
		fmt.Println(z)
	}(10, 20)

	add(20, 30) // calling the add function

}
