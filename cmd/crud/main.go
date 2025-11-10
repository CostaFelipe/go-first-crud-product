package main

import (
	"database/sql"

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
	db, err := Config()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	p := NewProduct("Laranja KG", 5.99)
	err = insertProduct(db, p)

	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name=?, price=? where id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = db.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}
