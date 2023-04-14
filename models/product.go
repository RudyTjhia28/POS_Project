package models

import "github.com/jinzhu/gorm"

// Product model
type Product struct {
	gorm.Model
	SKU         string  `json:"sku" gorm:"uniqueIndex"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       uint    `json:"stock"`
}
