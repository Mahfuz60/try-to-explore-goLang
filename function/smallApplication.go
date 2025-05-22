package main

import "fmt"

func main() {

	welcomeMassage()
	name := getUserName()
	income, cost := getUserSaving()
	saving := savingCalculation(income, cost)

	displayResult(saving, name)

}

func welcomeMassage() {
	fmt.Println("Welcome to our new Application")

}

func getUserName() string {
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)

	return name

}

func getUserSaving() (int, int) {
	var income int
	var cost int

	fmt.Println("Please enter your monthly income:")
	fmt.Scanln(&income)
	fmt.Println("Please enter your monthly cost:")
	fmt.Scanln(&cost)

	return income, cost

}

func savingCalculation(income int, cost int) int {
	saving := income - cost
	return saving
}

func displayResult(saving int, name string) {
	fmt.Println("Your monthly total saving is:", saving)
	fmt.Println("Thank you for using our application,", name, ",we are happy to help you!")
	fmt.Println("Have a nice day!")

}
