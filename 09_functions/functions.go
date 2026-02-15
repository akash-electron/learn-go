package main

import "fmt"

func main () {
	fmt.Println("Functions in GO")
	res := plus(2, 6)
	fmt.Println("sum : " , res)
	fmt.Println("returnMultipleValues in GO")

	fmt.Println(returnMultipleValues())
}

func plus(a, b int) int {
	// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a+b
}

func returnMultipleValues() (int ,int ) {
	return 3, 7
}
