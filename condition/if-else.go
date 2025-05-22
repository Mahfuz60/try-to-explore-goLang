package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age:")
	fmt.Scan(&age)

	if age > 18 {
		fmt.Println("You are eligible to vote")
	} else if age < 18 {
		fmt.Println("You are not eligible to vote")
	} else if age == 18 {
		fmt.Println("You are still a teenager,but you are eligible to vote after 18")
	}

}
