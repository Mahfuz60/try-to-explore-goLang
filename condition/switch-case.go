package main

import "fmt"

func main() {

	var a int = 3

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2, 3:
		fmt.Println("a is either 2 or 3")

	default:
		fmt.Println("a is neither 1,2 or 3")
	}

}
