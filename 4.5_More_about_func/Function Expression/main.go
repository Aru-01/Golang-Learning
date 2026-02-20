package main

import "fmt"

func main() {
	// Function Expression or Assign Function in variable

	add := func(a, b int) int {
		return a + b
	}

	res := add(12, 13)
	fmt.Println(res)
}
