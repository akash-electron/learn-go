package main

import "fmt"

func main() {
	// --- PART 1: IF / ELSE ---
	age := 18

	if age >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("Too young to vote.")
	}

	// --- PART 2: THE FOR LOOP ---
	// Go only has ONE loop: the 'for' loop.

	// A. Standard C-style loop
	fmt.Println("Counting to 3:")
	for i := 1; i <= 3; i++ {
		fmt.Println("Number:", i)
	}

	// B. Range-based loop (newest style in Go 1.22+)
	fmt.Println("Range 0 to 4 (skipping 2):")
	for i := range 5 {
		if i == 2 {
			continue // Skip number 2
		}
		fmt.Println("Index:", i)
	}

	// C. The "While" style loop
	counter := 10
	for counter > 0 {
		fmt.Print(counter, " ")
		counter -= 5
	}
	fmt.Println("Blast off!")
}
