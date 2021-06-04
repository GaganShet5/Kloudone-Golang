package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int{
		"Ram":     120000,
		"Hari":    150000,
		"Krishna": 900000,
	}
	employee := "Hari"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}
