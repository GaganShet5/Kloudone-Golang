package main

import "fmt"

func main() {

	arr := [7]string{"This", "hello", "the", "tasks",
		"of", "Go", "language"}

	fmt.Println("Array:", arr)

	myslice := arr[3:]

	fmt.Println("Slice:", myslice)
	fmt.Printf("Length of the slice: %d", len(myslice))
	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
}
