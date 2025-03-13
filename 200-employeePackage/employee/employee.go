package employee

import (
	"fmt"
)

type Employee struct {
	Name   string
	Salary int
	Stuff  int
}

func (e *Employee) SetupStuff(secret int) {
	e.Stuff = secret
}

func (e Employee) PrintEmployee() {
	fmt.Printf("Name: %s Salary: %d Stuff: %d\n", e.Name, e.Salary, e.Stuff)
}
