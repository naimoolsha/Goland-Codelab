package main

import "fmt"

// Variables Practice

func variablesPractice() {

	// Declaring variables
	// 1. Declaring a variable with a specific type and assigning a value to it.
	// for example:
	var message string = "Hello, World!"
	var city, country string = "New York", "USA" // multiple variable declaration with type inference
	var myName, myAge, amStudent = "John", 30, true // multiple variable declaration without type inference

	// 2. Declaring a variable without specifying a type, and letting the 
	// compiler infer the type based on the assigned value.
	// for example:
	var greeting = "Hello, Go!"

	// 3. Declaring a variable using the short variable declaration syntax (:=),
	//  which automatically infers the type based on the assigned value.
	// for example:
	name := "Alice"
	name1, isTrue := "Bob", true // multiple variable declaration using short variable declaration syntax

	// note: The short variable declaration syntax (:=) can only be used 
	// inside functions, not at the package level.

	// 4. Declaring a variable without assigning a value, which will result in
	// the variable being initialized with the zero value of its type.
	// for example:
	var uninitializedVar int
	

	// Print the variables
	fmt.Println("Declaring variables:")
	fmt.Println(message, greeting, name, uninitializedVar, myName, myAge, amStudent, city, country, name1, isTrue)


	/// -------------- Go Variable Types --------------
	// Go has several built-in variable types, including:
	// 1. String: Represents a sequence of characters.
	// 2. Integer: Represents whole numbers (e.g., int, int8, int16, int32, int64).
	// 3. Float: Represents floating-point numbers (e.g., float32, float64).
	// 4. Boolean: Represents true or false values (bool).
	// 5. Array: Represents a fixed-size collection of elements of the same type.
	// 6. Slice: Represents a dynamically-sized collection of elements of the same type.
	// 7. Map: Represents a collection of key-value pairs.
	// 8. Struct: Represents a composite data type that groups together related fields.


	// -------------- Variable Declaration Examples --------------
	// Here are some examples of variable declarations in Go:
	//
	// String variable
	var fullname string = "John Doe"

	// Integer variable
	var age int = 30

	// Float variable
	var height float64 = 5.9
	var weight float32 = 70.5

	// Boolean variable
	var isStudent bool = true

	// Array variable
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	// Slice variable
	var scores []int = []int{85, 90, 78, 92, 88}


	// Print the variables
	fmt.Println("Full Name:", fullname)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Weight:", weight)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Scores:", scores)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
}
