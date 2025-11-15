package db

import "database/sql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		return nil, err
	}

	return db, err
}
