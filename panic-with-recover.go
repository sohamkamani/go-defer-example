package main

import "fmt"

func main() {
	defer func() {
		// we can use the recover function inside the deferred function
		// to tell if it was called after `panic` was called
		if r := recover(); r != nil {
			fmt.Println("recovered from panic: ", r)
		}
	}()

	panic("something went wrong")
	fmt.Println("this should not be printed")
}

// this is printed before panic
// recovered from panic:  something went wrong
