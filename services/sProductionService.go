package services

import (
	"pos_project/models"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetProducts(ctx *gin.Context) (*[]models.Product, error)
	// GetProductById(id string) (*models.Product, error)
	// CreateProduct(*models.CreateProductRequest) (*models.Product, error)
	// UpdateProduct(*models.CreateProductRequest) (*models.Product, error)
}
