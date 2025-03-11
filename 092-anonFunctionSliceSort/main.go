package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name   string
	Salary int
}

func main() {
	staff := []Employee{
		{"Joe", 1000},
		{"Tom", 750},
		{"Harry", 1250},
	}

	sort.Slice(staff, func(i int, j int) bool { return staff[i].Salary < staff[j].Salary })

	fmt.Println(staff)
}
