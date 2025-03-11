package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	ID   int
	Name string
}

func main() {
	numberOfEmployees, spareCapacity := getEmployeeCountAndSpareCapcity()

	var employees []Employee = loadEmployees(numberOfEmployees, spareCapacity)

	fmt.Println("Number of employees:", len(employees))
	fmt.Println("Employees slice capacity:", cap(employees))

	fmt.Println(employees)
}

func getEmployeeCountAndSpareCapcity() (int, int) {
	return 13, 5 // obviously we would normally do something more useful here
}

func loadEmployees(employeeCnt int, spareCapacity int) []Employee {
	var result []Employee = make([]Employee, employeeCnt, (employeeCnt + spareCapacity))

	fmt.Println(employeeCnt)

	for i := 0; i < employeeCnt; i++ {

		result[i].ID = i
		result[i].Name = "Emp-" + strconv.Itoa(i)
	}

	return result
}
