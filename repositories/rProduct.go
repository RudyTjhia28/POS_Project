package repositories

import (
	"pos_project/models"

	"github.com/gin-gonic/gin"
)

type IProductRepository interface {
	GetProducts(ctx *gin.Context) (*[]models.Product, error)
	// GetProductById(ctx *gin.Context, id string) (models.Product, error)
	// CreateProduct(ctx *gin.Context, request models.CreateProductRequest) (models.Product, error)
	// UpdateProduct(ctx *gin.Context, request models.CreateProductRequest) (models.Product, error)
}
