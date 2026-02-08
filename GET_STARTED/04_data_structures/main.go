package main

import "fmt"

func main() {
	// --- PART 1: ARRAYS (Fixed Size) ---
	// Arrays have a size that NEVER changes.
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"
	fmt.Println("Array of Fruits:", fruits)

	// --- PART 2: SLICES (Dynamic Size) ---
	// Slices are like arrays but they can grow. These are used 99% of the time.
	vegetables := []string{"Carrot", "Potato"}
	vegetables = append(vegetables, "Spinach") // This adds an item!
	fmt.Println("Slice of Vegetables:", vegetables)

	// --- PART 3: 2D ARRAYS ---
	grid := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2D Grid:", grid)

	// --- PART 4: MAPS (Key-Value Pairs) ---
	// Like a dictionary.
	scores := make(map[string]int)
	scores["Akash"] = 95
	scores["Prachi"] = 100
	fmt.Println("Scores Map:", scores)
}
