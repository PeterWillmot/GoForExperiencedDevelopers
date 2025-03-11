package main

import (
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {
	tom := Employee{100, "Tom Jones", 20000}

	fmt.Println(tom.Name)
	fmt.Println(tom.Salary)
	tom.Salary = 30000
	fmt.Println(tom.Salary)
}
