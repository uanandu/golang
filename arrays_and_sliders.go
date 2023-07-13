package main

// Import - here we import the packages we need
// fmt - package for formatting text
// math - package for math operations
import (
	"fmt"
)

func main() {
	// Arrays
	myArray := [5]int{1, 2, 3, 4, 5} // Array of 5 integers

	// Slices
	mySlice := []int{1, 2, 3, 4, 5} // Slice of 5 integers

	fmt.Println("Here starts the array....")
	// When we loop through an array
	for index, value := range myArray {
		fmt.Println("Index:", index, "Value:", value)
	}
	fmt.Println("Here ends the array....")

	fmt.Println("#############################################")

	fmt.Println("Here starts the slice....")
	// When we loop through a slice
	for index, value := range mySlice {
		fmt.Println("Index:", index, "Value:", value)
	}
	fmt.Println("Here ends the slice....")
}
