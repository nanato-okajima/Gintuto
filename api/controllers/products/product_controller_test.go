package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"Gintuto/api/domain/products"
	"Gintuto/api/utils/errors"
)

func requestHandler(p interface{}) (*gin.Context, *httptest.ResponseRecorder) {
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
	p := products.Product{ID: 123, Name: "hoge taro"}
	c, res := requestHandler(p)

	CreateProduct(c)

	var product products.Product
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
	c, res := requestHandler(p)

	CreateProduct(c)

	var apiErr errors.ApiErr
	err := json.Unmarshal(res.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusBadRequest, res.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, "invalid json body", apiErr.Message)
	assert.EqualValues(t, 400, apiErr.Status)
	assert.EqualValues(t, "bad_request", apiErr.Error)
}
