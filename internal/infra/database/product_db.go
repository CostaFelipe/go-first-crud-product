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

func (p *Product) FindById(id string) (*product.Product, error) {
	stmt, err := p.DB.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product product.Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) Update(product *product.Product) error {
	stmt, err := p.DB.Prepare("update products set name=? price=? where id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Delete(id string) error {
	stmt, err := p.DB.Prepare("delete from products where id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
