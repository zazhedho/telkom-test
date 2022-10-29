package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm/models"
)

func TestGetAllProducts(t *testing.T) {
	repo := RepoProductMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Products{
		{ProductId: 1, Name: "test 1"},
		{ProductId: 2, Name: "test 2"},
	}

	repo.mock.On("FindAllProducts").Return(&dataMock, nil)
	data := service.GetAllProducts()

	result := data.Data.(*models.Products)
	for i, v := range *result {
		assert.Equal(t, dataMock[i].ProductId, v.ProductId, "expect id from data mock")
	}
}

func TestAddProduct(t *testing.T) {
	repo := RepoProductMock{mock.Mock{}}
	service := NewService(&repo)

	dataMock := models.Product{ProductId: 1, Name: "test 1"}

	repo.mock.On("SaveProduct", &dataMock).Return(&dataMock, nil)
	data := service.AddProduct(&dataMock)

	var expectId uint = 1
	result := data.Data.(*models.Product)

	assert.Equal(t, expectId, result.ProductId, "id must be 1")
}
