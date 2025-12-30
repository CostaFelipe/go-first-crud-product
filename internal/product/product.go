package product

import (
	"errors"

	"github.com/CostaFelipe/go-first-crud-productexample/pkg/id"
)

type Product struct {
	ID    id.ID   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var (
	errorIDIsRequired  = errors.New("ID is required")
	errInvalidID       = errors.New("ID is invalid")
	errorNameRequired  = errors.New("Name is required")
	errorPriceRequired = errors.New("Price is required")
)

func NewProduct(name string, price float64) (*Product, error) {
	id, err := id.NewID()
	if err != nil {
		return nil, err
	}
	product := &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}

	err = product.Validation()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validation() error {
	if p.ID.String() == "" {
		return errorIDIsRequired
	}
	if _, err := id.ParseID(p.ID.String()); err != nil {
		return errInvalidID
	}
	if p.Name == "" {
		return errorNameRequired
	}
	if p.Price == 0 {
		return errorPriceRequired
	}
	if p.Price < 0 {
		return errorPriceRequired
	}

	return nil
}
