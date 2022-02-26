package products

import (
	"Gintuto/api/datasources/mysql/products_db"
	"Gintuto/api/utils/errors"
	"fmt"
	"log"
)

var (
	productsDB = make(map[uint]*Product)
)

func (p *Product) Get() *errors.ApiErr {

	if err := products_db.Client.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	result := productsDB[p.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("product %d not found", p.ID))
	}

	p.ID = result.ID
	p.Name = result.Name
	p.Detail = result.Detail
	p.Price = result.Price
	p.Img = result.Img
	p.CreatedAt = result.CreatedAt
	p.UpdatedAt = result.UpdatedAt
	p.DeletedAt = result.DeletedAt

	return nil
}

func (p *Product) Save() *errors.ApiErr {
	current := productsDB[p.ID]
	if current != nil {
		if current.Name == p.Name {
			fmt.Println(current.Name)
			return errors.NewBadRequestError(fmt.Sprintf("name %s already registered", p.Name))
		}
		return errors.NewBadRequestError(fmt.Sprintf("product %d already exists", p.ID))
	}

	productsDB[p.ID] = p
	return nil
}
