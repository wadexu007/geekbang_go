package main

import (
	"fmt"
	"math"
)

// Main function
func main() {

	// Finding largest number
	// among the given numbers
	// Using Max() function
	res_1 := math.Max(0, -0)
	res_2 := math.Max(-100, 100)
	res_3 := math.Max(45.6, 8.9)
	res_4 := math.Max(math.NaN(), 67)
	res_5 := math.Abs(-3)

	// Displaying the result
	fmt.Printf("Result 1: %.1f", res_1)
	fmt.Printf("\nResult 2: %.1f", res_2)
	fmt.Printf("\nResult 3: %.1f", res_3)
	fmt.Printf("\nResult 4: %.1f", res_4)
	fmt.Printf("\nResult 5: %.0f \n", res_5)
}
