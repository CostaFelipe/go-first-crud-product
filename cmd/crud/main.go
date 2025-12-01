package main

import (
	"fmt"

	"github.com/CostaFelipe/go-first-crud-productexample/internal/db"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/product"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	p, err := product.NewProduct("Laranja KG", 5)

	//err = database.NewProduct(db).Create(p)
	fmt.Println(p.ID)
	if err != nil {
		panic(err)
	}
}
