package main

import "fmt"

func main() {
	// Anonymous Function
	// Immediately Invoked Function Expression
	// IIFE
	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20)

}
