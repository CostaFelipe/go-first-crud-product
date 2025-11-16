package main

import (
	"database/sql"

	"github.com/CostaFelipe/go-first-crud-productexample/internal/db"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/infra/database"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/product"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	p := product.NewProduct("Laranja KG", 5)

	err = database.NewProduct(db).Create(p)

	if err != nil {
		panic(err)
	}
}

func updateProduct(db *sql.DB, product *product.Product) error {
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

func removeProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id=?")
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

func selectProduct(db *sql.DB, id string) (*product.Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id=?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var p product.Product

	err = db.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func selectAllProducts(db *sql.DB) ([]product.Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []product.Product

	for rows.Next() {
		var p product.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
