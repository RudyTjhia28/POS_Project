package services

import "pos_project/models"

type ProductService interface {
	GetProducts() (*[]models.Product, error)
	GetProductById(id string) (*models.Product, error)
	CreateProduct(*models.CreateProductRequest) (*models.Product, error)
	UpdateProduct(*models.CreateProductRequest) (*models.Product, error)
}
