package main

import "fmt"

func sum(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	a := 4
	b := 6

	sum(a, b)
	sum(324, 45245)
}
