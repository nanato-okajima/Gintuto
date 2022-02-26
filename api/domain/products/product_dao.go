package products

import (
	"Gintuto/api/datasources/mysql/products_db"
	"Gintuto/api/utils/errors"
	"Gintuto/api/utils/mysqlutils"
)

var (
	productsDB = make(map[uint]*Product)
)

func (p *Product) Get() *errors.ApiErr {
	if result := products_db.Client.Where("id = ?", p.Model.ID).Find(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}

	return nil
}

func (p *Product) Save() *errors.ApiErr {
	if result := products_db.Client.Create(&p); result.Error != nil {
		mysqlutils.ParseError(result.Error)
	}

	return nil
}

func (p *Product) Update() *errors.ApiErr {
	if result := products_db.Client.Save(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

func (p *Product) PartialUpdate() *errors.ApiErr {
	if result := products_db.Client.Table("products").Where("id IN (?)", p.ID).Updates(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}

	return nil
}

func (p *Product) Delete() *errors.ApiErr {
	if result := products_db.Client.Delete(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}

	return nil
}
