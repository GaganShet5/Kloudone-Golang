package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int{
		"Ram":     1200000,
		"Hari":    1500000,
		"Krishna": 9000000,
	}
	employee := "Hari"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}
