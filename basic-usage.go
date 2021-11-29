package main

import "fmt"

func main() {
	// When we add `defer` before a function, that function is executed
	// after the surrounding function completes
	defer fmt.Println("this is printed once the function completes")

	fmt.Println("this is printed first")
	fmt.Println("this is printed second")
}
