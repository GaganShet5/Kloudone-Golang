package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to India")

	go func() {

		fmt.Println("Welcome to KloudOne")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("BeHappy")
}
