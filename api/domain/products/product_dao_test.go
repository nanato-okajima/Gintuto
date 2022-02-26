package products

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductGetNoError(t *testing.T) {
	p := Product{
		ID:        1,
		Name:      "hoge taro",
		Detail:    "hoge hoge",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	p.Save()
	newP := Product{ID: 1}

	result := newP.Get()

	assert.Nil(t, result)
	assert.EqualValues(t, p.Name, newP.Name)
	assert.EqualValues(t, p.Detail, newP.Detail)
	assert.EqualValues(t, p.Img, newP.Img)
}

func TestProductNotFound(t *testing.T) {
	p := Product{ID: 100}

	err := p.Get()

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "product 100 not found")
	assert.EqualValues(t, err.Status, 404)
	assert.EqualValues(t, err.Error, "not_found")

}

func TestProductSaveNoError(t *testing.T) {
	p := Product{
		ID:        2,
		Name:      "hoge taro",
		Detail:    "hoge hoge",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	err := p.Save()

	assert.Nil(t, err)
}

func TestProductSaveBadRequestErrorWithSameName(t *testing.T) {
	p := Product{
		ID:        1,
		Name:      "hoge taro",
		Detail:    "hoge hoge",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	p.Save()

	p2 := Product{
		ID:        1,
		Name:      "hoge taro",
		Detail:    "hoge hoge",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	err := p2.Save()

	assert.NotNil(t, err)
	assert.EqualValues(t, "name hoge taro already registered", err.Message)
	assert.EqualValues(t, 400, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestProductSaveBadRequestErrorWithSameID(t *testing.T) {
	p := Product{
		ID:        1,
		Name:      "hoge taro",
		Detail:    "hoge hoge",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	p.Save()

	p2 := Product{
		ID:        1,
		Name:      "hoge jiro",
		Detail:    "hoge hoge",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	err := p2.Save()

	assert.NotNil(t, err)
	assert.EqualValues(t, "product 1 already exists", err.Message)
	assert.EqualValues(t, 400, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
}
