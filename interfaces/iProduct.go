package interfaces

import "pos_project/models"

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	GetProductById(id string) (models.Product, error)
	CreateProduct(*models.CreateProductRequest) (*models.Product, error)
	UpdateProduct(*models.CreateProductRequest) (*models.Product, error)
}
