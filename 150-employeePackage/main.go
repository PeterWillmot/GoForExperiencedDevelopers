package main

import (
	"pkgdemo/employee"
)

// NOTE: we use GO MOD INIT command to create the go.mod file

// this is my main

func main() {
	// we should get an errors here because of unexported Employee code

	staff := []employee.Employee{
		{Name: "Joe", Salary: 1000, Stuff: 0},
		{Name: "Tom", Salary: 750, Stuff: 0},
		{Name: "Harry", Salary: 1250, Stuff: 0},
	}

	for i, e := range staff {
		e.SetupStuff(i)
		e.PrintEmployee()
	}
}
