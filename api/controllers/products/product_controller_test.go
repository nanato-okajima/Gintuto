package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	dproducts "Gintuto/api/domain/products"
	"Gintuto/api/utils/errors"
)

func getRequestHandler(id string) (*gin.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)

	param := gin.Param{Key: "product_id", Value: id}
	c.Params = gin.Params{param}

	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/products/:product_id",
		nil,
	)

	return c, res
}

func TestGetProductNoError(t *testing.T) {
	p := dproducts.Product{Model: gorm.Model{ID: 1}, Name: "hoge taro"}
	c, _ := postRequestHandler(p)
	CreateProduct(c)

	c2, res := getRequestHandler("1")

	GetProduct(c2)

	var product dproducts.Product
	err := json.Unmarshal(res.Body.Bytes(), &product)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.EqualValues(t, uint64(1), product.ID)
}

func TestGetProductWithInvalidID(t *testing.T) {
	c, res := getRequestHandler("a")

	GetProduct(c)

	var apiErr errors.ApiErr
	json.Unmarshal(res.Body.Bytes(), &apiErr)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, res.Code)
	assert.EqualValues(t, apiErr.Message, "product id should be a number")
	assert.EqualValues(t, apiErr.Status, 400)
	assert.EqualValues(t, apiErr.Error, "bad_request")
}

func TestGetProductWithNoProduct(t *testing.T) {
	c, res := getRequestHandler("10000")

	GetProduct(c)

	var apiErr errors.ApiErr
	json.Unmarshal(res.Body.Bytes(), &apiErr)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusNotFound, res.Code)
	assert.EqualValues(t, apiErr.Message, "product 10000 not found")
	assert.EqualValues(t, apiErr.Status, 404)
	assert.EqualValues(t, apiErr.Error, "not_found")
}

func postRequestHandler(p interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	byteProduct, _ := json.Marshal(p)
	c, _ := gin.CreateTestContext(res)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/products",
		bytes.NewBuffer(byteProduct),
	)

	return c, res
}

func TestCreateProductNoError(t *testing.T) {
	p := dproducts.Product{Model: gorm.Model{ID: 123}, Name: "hoge taro"}
	c, res := postRequestHandler(p)

	CreateProduct(c)

	var product dproducts.Product
	err := json.Unmarshal(res.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusCreated, res.Code)
	assert.Nil(t, err)
	fmt.Println(product)
	assert.EqualValues(t, uint64(123), product.ID)
}

func TestCreateProductWith404Error(t *testing.T) {
	type demiProduct struct {
		ID   string `json"id"`
		Name string `json:"name"`
	}

	p := demiProduct{ID: "123", Name: "hoge taro"}
	c, res := postRequestHandler(p)

	CreateProduct(c)

	var apiErr errors.ApiErr
	err := json.Unmarshal(res.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusBadRequest, res.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, "invalid json body", apiErr.Message)
	assert.EqualValues(t, 400, apiErr.Status)
	assert.EqualValues(t, "bad_request", apiErr.Error)
}
