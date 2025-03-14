package demodb

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

var pgConnection *sql.DB // NOTE: Name starts lower case - not exported!

func GetConnection() *sql.DB {
	return pgConnection
}

func Connect() (bool, error) {
	pgConnectionString := "postgres://postgres:pgpassword@localhost:5432/demodb?sslmode=disable"

	var err error

	pgConnection, err = sql.Open("postgres", pgConnectionString)

	if err != nil {
		log.Println("OPEN:", err)
		return false, errors.New("Cant connect to PosgresDB: ")
	}

	if err = pgConnection.Ping(); err != nil {
		log.Println("PING:", err)
		return false, errors.New("Cant PING PosgresDB")
	}

	return true, nil
}

func Disconnect() {
	if pgConnection != nil {
		pgConnection.Close()
	}

	pgConnection = nil
}
