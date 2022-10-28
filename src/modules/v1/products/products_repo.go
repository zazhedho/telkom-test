package products

import (
	"errors"

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

	if result.Error != nil {
		return nil, errors.New("data tidak dapat ditampilkan")
	}
	return &data, nil
}

func (re *product_repo) SaveProduct(data *models.Product) (*models.Product, error) {
	result := re.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("gagal menambah data")
	}
	return data, nil
}

func (re *product_repo) ChangeProduct(id int, data *models.Product) (*models.Product, error) {
	result := re.db.Model(&data).Where("product_id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, errors.New("gagal update data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data product tidak ditemukan")
	}

	return data, nil
}

func (re *product_repo) RemoveProduct(kode string, data *models.Product) (*models.Product, error) {
	result := re.db.Where("LOWER(kode_product) = ?", kode).Delete(data)

	if result.Error != nil {
		return nil, errors.New("gagal menghapus data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data product tidak ditemukan")
	}

	return nil, nil
}

func (re *product_repo) SortByName(name string, data *models.Products) (*models.Products, error) {
	result := re.db.Order("created_at desc").Where("LOWER(name) LIKE ?", "%"+name+"%").Find(&data)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("data product tidak ditemukan")
	}

	return data, nil
}
