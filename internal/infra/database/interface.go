package database

import "github.com/CostaFelipe/go-first-crud-productexample/internal/product"

type ProductInterface interface {
	Create(product *product.Product) error
	FindAll() ([]product.Product, error)
	FindById(id string) (*product.Product, error)
	Update(product *product.Product) error
	Delete(id string) error
}
