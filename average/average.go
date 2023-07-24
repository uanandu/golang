package main

import (
	"fmt"
)

func main() {
	// This program finds the average of the numbers given

	// array
	numbers := [3]float64{71.8, 56.2, 89.5}

	var sum float64 = 0

	// here we dont consider the index
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}
