package main

import "fmt"

var i int

func count() int {
	i++
	return i
}

func main() {

	defer fmt.Println("deferred count:", count())

	fmt.Println("current count:", count())
}
