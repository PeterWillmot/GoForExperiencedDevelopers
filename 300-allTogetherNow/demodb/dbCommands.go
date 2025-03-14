package demodb

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DropEmployeeTable() {

	fmt.Println("db.DropEmployeeTable")

	if pgConnection == nil {
		log.Fatal("No pgConnection")
	}

	sql := `DROP TABLE IF EXISTS employee`

	_, err := pgConnection.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateEmployeeTable() {
	if pgConnection == nil {
		log.Fatal("No pgConnection")
	}

	sql := `DROP INDEX IF EXISTS ix_employee_name`

	_, err := pgConnection.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	sql = `CREATE TABLE IF NOT EXISTS employee
				(
					id SERIAL PRIMARY KEY
					,name VARCHAR(100) NOT NULL
					,salary NUMERIC(8,2)
				)`

	_, err = pgConnection.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	sql = `CREATE UNIQUE INDEX ix_employee_name ON employee(name)`

	_, err = pgConnection.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}
}
