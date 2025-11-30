package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CostaFelipe/go-first-crud-productexample/internal/dto"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/infra/database"
	"github.com/CostaFelipe/go-first-crud-productexample/internal/product"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProductHandle(w http.ResponseWriter, r *http.Request) {
	var prod dto.ProductInput
	err := json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := product.NewProduct(prod.Name, prod.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
