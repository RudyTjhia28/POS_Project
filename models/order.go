package models

import "github.com/jinzhu/gorm"

// Order model
type Order struct {
	gorm.Model
	OrderNumber  string      `json:"order_number" gorm:"uniqueIndex"`
	CustomerName string      `json:"customer_name"`
	TotalAmount  float64     `json:"total_amount"`
	OrderItems   []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

// OrderItem model
type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}
