package main

import "fmt"

// Print func
/*
func sum(a int, b int) {
	fmt.Println(a + b)
}
*/

// return func
func sum_and_mult(a int, b int) (int, int) {
	sum := a + b
	mult := a * b
	return sum, mult
}

func main() {
	a := 4
	b := 6

	x, y := sum_and_mult(a, b)
	fmt.Println(x, y)
	x, y = sum_and_mult(324, 45245)
	fmt.Println(x, y)
}
