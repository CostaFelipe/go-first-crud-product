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

func ConfigDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goproduct")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := ConfigDB()
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

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id=?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var p Product

	err = db.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil

}
