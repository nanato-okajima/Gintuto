package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dproducts "Gintuto/api/domain/products"
	"Gintuto/api/services"
	"Gintuto/api/utils/errors"
)

func GetProduct(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("product id should be a number")
		c.JSON(err.Status, err)
		return
	}

	product, getErr := services.GetProduct(uint(productID))
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product dproducts.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	newProduct, saveErr := services.CreateProduct(product)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}
