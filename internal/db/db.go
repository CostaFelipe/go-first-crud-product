package db

import "database/sql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/bancodeteste")
	if err != nil {
		return nil, err
	}
	return db, err
}
