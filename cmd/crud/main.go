package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    NewID().String(),
		Name:  name,
		Price: price,
	}
}

func Config() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goproduct")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	fmt.Println("Hello, world!")
}
