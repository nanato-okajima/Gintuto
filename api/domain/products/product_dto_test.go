package products

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestProductValidateNoError(t *testing.T) {
	p := Product{Model: gorm.Model{ID: 123}, Name: "hoge taro"}

	err := p.Validate()

	assert.Nil(t, err)
}

func TestProcutValidateBadRequestError(t *testing.T) {
	p := Product{Model: gorm.Model{ID: 123}}

	err := p.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid product name", err.Message)
	assert.EqualValues(t, 400, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
}
