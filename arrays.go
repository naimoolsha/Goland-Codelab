package main

import "fmt"

/// learning about arrays in Go
/// 1. Declaring an array with a specific size and type, and assigning values to it.
func declaringArrays() {
	// ways of Declaring an array
	// 1. Declaring an array with a specific size and type, and assigning values to it.
	// for example:
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	// 2. Declaring an array without specifying the size, and letting the compiler
	// infer the size based on the number of assigned values.
	// for example:
	var fruits = [...]string{"Apple", "Banana", "Cherry"}

	// 3. Declaring an array using the short variable declaration syntax (:=),
	// which automatically infers the size and type based on the assigned values.
	// for example:
	cities := [...]string{"New York", "Los Angeles", "Chicago"}

	// Print the arrays
	fmt.Println("Declaring arrays:")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Fruits:", fruits)
	fmt.Println("Cities:", cities)
}

/// 2. Accessing and Modifying Array Elements
func accessingAndModifyingArrays() {
	// Declaring an array of integers
	numbers := [5]int{10, 20, 30, 40, 50}

	// Accessing array elements using their index
	fmt.Println("Accessing array elements:")
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])
	fmt.Println("Third element:", numbers[2])

	// Modifying array elements by assigning new values to specific indices
	numbers[1] = 25 // Changing the second element from 20 to 25
	numbers[3] = 45 // Changing the fourth element from 40 to 45

	// Print the modified array
	fmt.Println("Modified array:", numbers)
}


