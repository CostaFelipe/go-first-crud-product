package handlers

import "github.com/CostaFelipe/go-first-crud-productexample/internal/infra/database"

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}
