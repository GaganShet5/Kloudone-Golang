package main

import (
	"fmt"
)

func main() {
	b := 24
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}
