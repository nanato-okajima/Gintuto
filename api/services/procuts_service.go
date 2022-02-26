package services

import (
	"Gintuto/api/domain/products"
	"Gintuto/api/utils/errors"
)

func GetProduct(productID uint64) (*products.Product, *errors.ApiErr) {
	p := &products.Product{ID: productID}
	if err := p.Get(); err != nil {
		return nil, err
	}

	return p, nil
}

func CreateProduct(product products.Product) (*products.Product, *errors.ApiErr) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	if err := product.Save(); err != nil {
		return nil, err
	}

	return &product, nil
}
