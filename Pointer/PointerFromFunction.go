package main

import (
	"fmt"
)

func hello() *int {
	i := 6
	return &i
}
func main() {
	d := hello()
	fmt.Println("Value of d", *d)
}
