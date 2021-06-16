// Pointer variables
package main

import "fmt"

func main() {

	// using := operator to declare
	// and initialize the variable
	y := 25

	// taking a pointer variable using
	// := by assigning it with the
	// address of variable y
	p := &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
}
