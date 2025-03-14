package employee

import (
	"demoapp/demodb"
	"fmt"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

func (e *Employee) Init(name string, salary float64) {
	e.ID = 0
	e.Name = name
	e.Salary = salary
}

func (e *Employee) Print() {
	fmt.Printf("ID: %d Name: %s Salary: %.2f\n", e.ID, e.Name, e.Salary)
}

func (e *Employee) GetByName() bool {
	sql := `SELECT id, salary FROM employee WHERE name = $1`

	demodb.GetConnection().QueryRow(sql, e.Name).Scan(&e.ID, &e.Salary)

	if e.ID > 0 {
		fmt.Println("FOUND", e.Name, "AS:", e.ID)
		return true
	} else {
		return false
	}
}

func (e *Employee) GetByID() bool {
	if e.ID > 0 {
		sql := `SELECT name, salary FROM employee WHERE id = $1`

		demodb.GetConnection().QueryRow(sql, e.ID).Scan(&e.Name, &e.Salary)

		if e.Name != "" {
			fmt.Println("FOUND", e.ID, "AS:", e.Name)
			return true
		} else {
			return false
		}
	}

	return false
}

func (e *Employee) Save() (bool, error) {

	if e.ID == 0 && !e.GetByName() {
		sql := `INSERT INTO employee(name, salary) VALUES($1, $2) RETURNING id`

		err := demodb.GetConnection().QueryRow(sql, e.Name, e.Salary).Scan(&e.ID)

		if err != nil {
			return false, err
		}

		fmt.Println("INSERTED", e.Name, "AS:", e.ID)

		return true, nil
	}

	return false, nil
}
