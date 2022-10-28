package models

import "time"

type Product struct {
	ProductId   uint      `gorm:"primaryKey" json:"id,omitempty"`
	KodeProduct string    `gorm:"not null" json:"kode_product,omitempty"`
	Name        string    `gorm:"not null" json:"name,omitempty"`
	Kuantitas   int       `gorm:"not null" json:"kuantitas,omitempty"`
	CreatedAt   time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Products []Product
