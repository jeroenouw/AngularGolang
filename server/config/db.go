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
	DB, err = sql.Open("postgres", "postgres://bond:password@localhost/ngxgo?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
