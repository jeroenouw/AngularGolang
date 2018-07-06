package data

import (
	"errors"
	"net/http"

	"github.com/jeroenouw/AngularGolang/server/config"
)

type User struct {
	ID   int
	name string
	city string
	age  int
}

func AllUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.name, &u.city, &u.age) // order matters
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func SingleUser(r *http.Request) (User, error) {
	u := User{}
	id := r.FormValue("ID")
	if id == "" {
		return u, errors.New("400. Bad Request.")
	}

	row := config.DB.QueryRow("SELECT * FROM users WHERE ID = $1", id)

	err := row.Scan(&u.ID, &u.name, &u.city, &u.age)
	if err != nil {
		return u, err
	}

	return u, nil
}
