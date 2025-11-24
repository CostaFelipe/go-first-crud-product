package product

import "github.com/CostaFelipe/go-first-crud-productexample/pkg/id"

type Product struct {
	ID    id.ID
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    id.NewID(),
		Name:  name,
		Price: price,
	}
}
