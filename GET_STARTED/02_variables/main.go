package main

import "fmt"

func main() {
	// 1. Standard Variable Declaration
	// Use 'var' when you want to be very specific about the type.
	var name string = "Akash"
	var health int = 100
	fmt.Println("Name:", name, "Health:", health)

	// 2. Short Variable Declaration (The shorthand)
	// Use ':=' inside functions. Go figures out the type for you.
	level := 1
	isAlive := true
	fmt.Println("Level:", level, "Alive:", isAlive)

	// 3. Constants
	// Use 'const' for things that NEVER change.
	const Gravity = 9.8
	fmt.Println("Gravity on Earth:", Gravity)

	// 4. Grouped Constants
	// This is a clean way to define multiple related constants.
	const (
		StatusActive   = 0
		StatusInactive = 1
		StatusPending  = 2
	)
	fmt.Println("Status Codes:", StatusActive, StatusInactive, StatusPending)
}
