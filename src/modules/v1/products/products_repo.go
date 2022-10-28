package products

import (
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm/models"
	"gorm.io/gorm"
)

type product_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *product_repo {
	return &product_repo{db}
}

func (re *product_repo) FindAllProducts() (*models.Products, error) {
	var data models.Products
	result := re.db.Order("created_at desc").Find(&data)
}
