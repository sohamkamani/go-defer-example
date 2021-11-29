package main

import "fmt"

func main() {
	myFunction()
}

func myFunction() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")

	fmt.Println("starting myFuntion...")
	// ...
}
