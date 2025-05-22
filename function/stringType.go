package main

import "fmt"

func main() {

	printSomethings()
	sayHello("Mahfuz")

}

func printSomethings() {
	fmt.Println("Education must be free")
}

func sayHello(name string) {
	fmt.Println("Welcome to our new golang course, ", name)
}
