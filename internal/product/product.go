package product

import "github.com/CostaFelipe/go-first-crud-productexample/pkg/id"

type Product struct {
	ID    id.ID
	Name  string
	Price int
}

func NewProduct(name string, price int) *Product {
	return &Product{
		ID:    id.NewID(),
		Name:  name,
		Price: price,
	}
}
