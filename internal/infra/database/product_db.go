package database

import (
	"database/sql"

	"github.com/CostaFelipe/go-first-crud-productexample/internal/product"
)

type Product struct {
	DB *sql.DB
}

func NewProduct(db *sql.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *product.Product) error {
	stmt, err := p.DB.Prepare("insert into products(id, name, price) values(?, ?, ?)")
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
