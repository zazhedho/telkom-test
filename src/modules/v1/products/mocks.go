package products

import (
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm/models"
)

type RepoProductMock struct {
	mock mock.Mock
}

func (m *RepoProductMock) FindAllProducts() (*models.Products, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Products), nil
}

func (m *RepoProductMock) SaveProduct(data *models.Product) (*models.Product, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.Product), nil
}

func (m *RepoProductMock) ChangeProduct(id int, data *models.Product) (*models.Product, error) {
	args := m.mock.Called(id, data)
	return args.Get(0).(*models.Product), nil
}

func (m *RepoProductMock) RemoveProduct(kode string, data *models.Product) (*models.Product, error) {
	args := m.mock.Called(kode, data)
	return args.Get(0).(*models.Product), nil
}

func (m *RepoProductMock) SortByName(name string, data *models.Products) (*models.Products, error) {
	args := m.mock.Called(name, data)
	return args.Get(0).(*models.Products), nil
}

func (m *RepoProductMock) SortByQty(qty int, data *models.Products) (*models.Products, error) {
	args := m.mock.Called(qty, data)
	return args.Get(0).(*models.Products), nil
}
