package main

import (
	"net/http"

	"github.com/CostaFelipe/go-first-crud-productexample/internal/db"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/infra/database"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/infra/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r.Post("/products/create", productHandler.CreateProductHandler)
	r.Get("/products/id", productHandler.GetProductHandler)
	r.Get("/products", productHandler.GetProducts)
	r.Delete("/products/id", productHandler.DeleteProduct)

	http.ListenAndServe(":3000", r)
}
