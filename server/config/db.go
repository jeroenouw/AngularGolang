package config

import (
	"database/sql"
	"fmt"
	// postgres
	_ "github.com/lib/pq"
)

// DB ...
var DB *sql.DB

func init() {
	var err error
	connStr := "user=jeroenouw dbname=ngxgo sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
