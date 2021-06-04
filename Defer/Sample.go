package main

import "fmt"

func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}
func show() {
	fmt.Println("Good Morning")
}
func main() {
	mul(13, 27)
	defer mul(23, 56)
	show()
}
