package services

import (
	"pos_project/models"
	"pos_project/repositories"

	"github.com/gin-gonic/gin"
)

type ProductServices struct {
	productRepo *repositories.IProductRepository
}

func NewProductService(repositories *repositories.IProductRepository) *repositories.IProductRepository {
	return repositories
}

func (ps *ProductServices) GetProducts(c *gin.Context) (*[]models.Product, error) {
	products, err := ps.productRepo.GetProducts(c)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// func (ps *ProductServices) GetProductById(id string) (*models.Product, error) {
// 	product, err := ps.productRepo.GetProductById(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &product, nil
// }

// func (ps *ProductServices) CreateProduct(input *models.CreateProductRequest) (*models.Product, error) {
// 	product := models.Product{Name: input.Name, Price: input.Price, Quantity: input.Quantity}
// 	result, err := ps.productRepo.CreateProduct(models.CreateProductRequest(product))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &result, nil
// }

// func (ps *ProductServices) UpdateProduct(input *models.CreateProductRequest) (*models.Product, error) {
// 	result, err := ps.productRepo.UpdateProduct(*input)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }
