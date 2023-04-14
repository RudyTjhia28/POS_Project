package models

import "github.com/jinzhu/gorm"

// Product model
type Product struct {
	gorm.Model
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
}
