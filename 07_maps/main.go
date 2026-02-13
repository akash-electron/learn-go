package main

import "fmt"

func main() {
	// Maps are key-value pair
	// maps => hash , objects , hash tables
	// Syntax: map[keyType]valueType
	m := make(map[string] int)
	fmt.Println(m)
    // setting key value pairs
    m["akash"] = 22
	m["prachiee"] = 21
    fmt.Println(m)

	// deleting a key value pair
	// if any key value pair is not present then it returns zero
	delete(m, "prachiee")
    fmt.Println(m)
}
