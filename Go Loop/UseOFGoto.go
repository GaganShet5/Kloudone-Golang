package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 10

	/* do loop execution */
LOOP:
	for a < 25 {
		if a == 15 {

			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}
