package main

import (
	"fmt"
)

func print_hello() {
	fmt.Println("hello")
	print_hello()
}
func main() {
	print_hello()
}
