package main

import (
	"fmt"
)

type Department struct {
	ID   int
	Name string
}

type Employee struct {
	ID         int
	Name       string
	Salary     float64
	Department *Department
	Manager    *Employee
}

func (e Employee) printEmployee() {
	fmt.Printf("(%02d) %s Manager: %s Salary: %.2f\n", e.ID, e.Name, e.Manager.Name, e.Salary)
}

func (e *Employee) adjustSalary(newSalary float64) {
	e.Salary = newSalary
}

func main() {
	management := Department{1, "Management"}
	research := Department{2, "Research"}

	tom := Employee{100, "Tome Jones", 20000, &management, nil}
	joe := Employee{100, "Joe Bloggs", 10000, &research, &tom}

	joe.printEmployee()
	joe.adjustSalary(15000)
	joe.printEmployee()
}
