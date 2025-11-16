package database

import "github.com/CostaFelipe/go-first-crud-productexample/internal/product"

type ProductInterface interface {
	Create(product *product.Product) error
}
