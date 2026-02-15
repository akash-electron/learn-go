package main

import (
	"fmt"
)

// range used for iterating over data structures

func main() {
	// iterate over slices or array
	nums := []int {1, 4,5 ,6 ,7}
	for i , num := range nums{
		fmt.Println(i , num)
	}
	// sum the numbers in a slice

	// range on arrays and slices provides both the index and value for each entry. 

	sum :=0
	for _ , num := range nums{
		sum = sum + num
	}

	fmt.Println("sum : " ,sum)


	// range on map iterates over key/value pairs.
	// range can also iterate over just the keys of a map.
	fmt.Println("Iterating over maps")

	m := map[string]int {"age" : 23 , "YOE" : 34}
	
	for k , v := range m {
		fmt.Println(k ,  "->" ,v)
	}

	for _ , k := range m {
		fmt.Println(k, "->" )
	}

}