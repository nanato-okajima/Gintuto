package products

import (
	"Gintuto/api/datasources/mysql/products_db"
	"Gintuto/api/utils/errors"
	"fmt"
)

var (
	productsDB = make(map[uint]*Product)
)

func (p *Product) Get() *errors.ApiErr {
	if result := products_db.Client.Where("id = ?", p.Model.ID).Find(&p); result.Error != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get product: %s", result.GetErrors()))
	}

	return nil
}

func (p *Product) Save() *errors.ApiErr {
	if result := products_db.Client.Create(&p); result.Error != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save product: %s", result.GetErrors()),
		)
	}

	return nil
}
