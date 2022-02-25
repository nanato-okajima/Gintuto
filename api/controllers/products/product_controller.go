package products

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Gintuto/api/domain/products"
	"Gintuto/api/utils/errors"
)

func GetProduct(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet implement!")
}

func CreateProduct(c *gin.Context) {
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.JSON(http.StatusCreated, product)
}
