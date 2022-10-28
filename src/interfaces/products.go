package interfaces

import "task6-api/src/database/orm/models"

type ProductRepo interface {
	FindAllProducts() (*models.Products, error)
}
