package interfaces

import (
	"pos_project/models"

	"github.com/gin-gonic/gin"
)

type ProductRepository interface {
	GetProducts(ctx *gin.Context) (*[]models.Product, error)
	// GetProductById(ctx *gin.Context, id string) (*models.Product, error)
	// CreateProduct(ctx *gin.Context, req *models.CreateProductRequest) (*models.Product, error)
	// UpdateProduct(ctx *gin.Context, req *models.CreateProductRequest) (*models.Product, error)
}
