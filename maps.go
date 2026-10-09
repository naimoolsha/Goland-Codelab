package main

import "fmt"

// learning about maps in Go
func mapsPractice() {
	// Declaring a map with string keys and int values
	ages := make(map[string]int)

	// Adding key-value pairs to the map
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	// Accessing values from the map using keys
	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Bob's age:", ages["Bob"])
	fmt.Println("Charlie's age:", ages["Charlie"])

	// Checking if a key exists in the map
	if age, exists := ages["David"]; exists {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David's age not found")
	}

	// Deleting a key-value pair from the map
	delete(ages, "Bob")

	// Iterating over the map using a for loop
	fmt.Println("Ages after deleting Bob:")
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}
}
