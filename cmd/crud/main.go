package main

import (
	"net/http"

	"github.com/CostaFelipe/go-first-crud-productexample/internal/db"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/infra/database"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/infra/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := http.NewServeMux()

	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	mux.HandleFunc("/products", productHandler.CreateProductHandle)

	http.ListenAndServe(":3000", mux)
}
