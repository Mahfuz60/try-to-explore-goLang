package main

import "fmt"

//global scope
var (
	e = 20
	f = 10
	g = 30
)

func main() {
	//local scope
	var age int
	fmt.Println("Enter you age:")
	fmt.Scanln(&age)

	//block scope{}
	if age >= 18 {
		p := 5
		fmt.Println("You are adult")
		fmt.Println("you are't access", p, "website")

	} else {
		fmt.Println("You are't adult,still a child")
	}
}
