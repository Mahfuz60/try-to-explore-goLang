package main

import "fmt"

var a int = 20

func main() {
	age := 25
	if age > 20 {
		a := 35 //This 'a' shadows the global variable 'a'
		fmt.Println("Inside if block a:", a)
	}

	fmt.Println("Outside block a:", a)
}
