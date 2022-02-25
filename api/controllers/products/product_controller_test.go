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

	"Gintuto/domain/products"
)

func TestCreateProduct(t *testing.T) {
	p := products.Product{ID: 123, Name: "hoge taro"}
	byteProduct, _ := json.Marshal(p)
	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/products",
		bytes.NewBuffer(byteProduct),
	)

	CreateProduct(c)

	var product products.Product
	err := json.Unmarshal(res.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusOK, res.Code)
	assert.Nil(t, err)
	fmt.Println(product)
	assert.EqualValues(t, uint64(123), product.ID)

}
