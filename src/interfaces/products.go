package interfaces

import (
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm/models"
	"github.com/zazhedho/telkom-test/task6-api/src/helpers"
)

type ProductRepo interface {
	FindAllProducts() (*models.Products, error)
	SaveProduct(data *models.Product) (*models.Product, error)
	ChangeProduct(id int, data *models.Product) (*models.Product, error)
	RemoveProduct(kode string, data *models.Product) (*models.Product, error)
	SortByName(name string, data *models.Products) (*models.Products, error)
}

type ProductService interface {
	GetAllProducts() *helpers.Response
	AddProduct(data *models.Product) *helpers.Response
	UpdateProduct(id int, data *models.Product) *helpers.Response
	DeleteProduct(kode string, data *models.Product) *helpers.Response
	SortByName(name string, data *models.Products) *helpers.Response
}
