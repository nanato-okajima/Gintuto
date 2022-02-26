package services

import (
	"github.com/jinzhu/gorm"

	"Gintuto/api/domain/products"
	"Gintuto/api/utils/errors"
)

func GetProduct(productID uint) (*products.Product, *errors.ApiErr) {
	p := &products.Product{Model: gorm.Model{ID: productID}}
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

func UpdateProduct(isPartial bool, product products.Product) (*products.Product, *errors.ApiErr) {
	current, err := GetProduct(product.ID)
	if err = current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if product.Name != "" {
			current.Name = product.Name
		}
		if product.Detail != "" {
			current.Detail = product.Detail
		}
		if product.Price != 0 {
			current.Price = product.Price
		}
		if product.Img != nil {
			current.Img = product.Img
		}
		if err := current.PartialUpdate(); err != nil {
			return nil, err
		}
	} else {
		current.Name = product.Name
		current.Detail = product.Detail
		current.Price = product.Price
		current.Img = product.Img
		if err := current.Update(); err != nil {
			return nil, err
		}
	}

	return current, nil
}

func DeleteProduct(productID uint) *errors.ApiErr {
	product := &products.Product{Model: gorm.Model{ID: productID}}
	return product.Delete()
}
