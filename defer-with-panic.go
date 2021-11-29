package main

import "fmt"

func main() {
	// When we add `defer` before a function, that function is executed
	// after the surrounding function completes
	defer fmt.Println("this is printed once the function panics")

	fmt.Println("this is printed before panic")
	panic("something went wrong")
	fmt.Println("this should not be printed")
}

// this is printed before panic
// this is printed once the function panics
// panic: something went wrong

// goroutine 1 [running]:
// main.main()
//         /Users/soham/src/github.com/sohamkamani/go-defer-example/defer-with-panic.go:11 +0xe4
// exit status 2
