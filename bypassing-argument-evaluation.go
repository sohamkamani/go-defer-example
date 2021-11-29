package main

import "fmt"

var i int

func count() int {
	i++
	return i
}

func main() {

	defer printCount()

	fmt.Println("current count:", count())
}

func printCount() {
	fmt.Println("count:", count())
}
