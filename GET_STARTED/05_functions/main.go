package main

import "fmt"

// add is a basic function that takes two integers and returns their sum.
func add(a int, b int) int {
	return a + b
}

// greet is a function that doesn't return anything (void).
func greet(name string) {
	fmt.Println("Hello there,", name)
}

// divide is a cool Go function that returns TWO things. 
// In Go, we often return the result AND an error message.
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Cannot divide by zero!"
	}
	return a / b, ""
}

func main() {
	// 1. Call a simple function
	result := add(10, 20)
	fmt.Println("10 + 20 =", result)

	// 2. Call a void function
	greet("Akash")

	// 3. Call a function with multiple returns
	val, err := divide(10, 2)
	fmt.Println("10 / 2 =", val, "| Error:", err)

	val2, err2 := divide(10, 0)
	fmt.Println("10 / 0 =", val2, "| Error:", err2)
}
