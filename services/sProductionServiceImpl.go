package services

import (
	"pos_project/models"
	"pos_project/repositories"

	"gorm.io/gorm"
)

type productService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(db *gorm.DB) *productService {
	productRepo := repositories.NewProductRepository(db)
	return &productService{productRepo}
}

func (ps *productService) GetProducts() (*[]models.Product, error) {
	products, err := ps.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (ps *productService) GetProductById(id string) (*models.Product, error) {
	product, err := ps.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (ps *productService) CreateProduct(input *models.CreateProductRequest) (*models.Product, error) {
	product := models.Product{Name: input.Name, Price: input.Price, Quantity: input.Quantity}
	result, err := ps.productRepo.CreateProduct(models.CreateProductRequest(product))
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ps *productService) UpdateProduct(input *models.CreateProductRequest) (*models.Product, error) {
	result, err := ps.productRepo.UpdateProduct(*input)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
