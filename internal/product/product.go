package product

import "github.com/CostaFelipe/go-first-crud-productexample/pkg/id"

type Product struct {
	ID    id.ID   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) (*Product, error) {
	return &Product{
		ID:    id.NewID(),
		Name:  name,
		Price: price,
	}, nil
}
