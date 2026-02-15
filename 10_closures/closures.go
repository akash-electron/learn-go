package main
import "fmt"

func main(){
	// closures are functions that reference variables from outside their body.
	// the function returned by firstFunction is a closure that references the variable i from the outer function.
	closure := firstFunction()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())

	newClosure :=firstFunction()
	fmt.Println(newClosure())
	fmt.Println(newClosure())

}
//firstFunction returns a function that remembers a variable (i) even after firstFunction has finished executing.
func firstFunction() func ()int {
 i:=0
  return func () int   {
	i++
	return i
  }
}