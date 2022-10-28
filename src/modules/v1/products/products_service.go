package products

import (
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm/models"
	"github.com/zazhedho/telkom-test/task6-api/src/helpers"
	"github.com/zazhedho/telkom-test/task6-api/src/interfaces"
)

type product_service struct {
	repo interfaces.ProductRepo
}

func NewService(repo interfaces.ProductRepo) *product_service {
	return &product_service{repo}
}

func (s *product_service) GetAllProducts() *helpers.Response {
	result, err := s.repo.FindAllProducts()
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}
	return helpers.New(result, 200, false)
}

func (s *product_service) AddProduct(data *models.Product) *helpers.Response {
	result, err := s.repo.SaveProduct(data)
	if err != nil {
		return helpers.New(err.Error(), 400, true)
	}
	return helpers.New(result, 201, false)
}

func (s *product_service) UpdateProduct(id int, data *models.Product) *helpers.Response {
	result, err := s.repo.ChangeProduct(id, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *product_service) DeleteProduct(kode string, data *models.Product) *helpers.Response {
	result, err := s.repo.RemoveProduct(kode, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}

func (s *product_service) SortByName(name string, data *models.Products) *helpers.Response {
	result, err := s.repo.SortByName(name, data)
	if err != nil {
		return helpers.New(err.Error(), 404, true)
	}

	return helpers.New(result, 200, false)
}
