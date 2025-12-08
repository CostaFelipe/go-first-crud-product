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

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
