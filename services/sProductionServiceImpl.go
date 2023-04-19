package services

import (
	"errors"
	"pos_project/models"

	"gorm.io/gorm"
)

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *productService {
	return &productService{db}
}

func (ps *productService) GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := ps.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *productService) GetProductById(id string) (models.Product, error) {
	var product models.Product
	if err := ps.db.First(&product, id).Error; err != nil {
		return models.Product{}, errors.New("Record not found")
	}

	return product, nil
}

func (ps *productService) CreateProduct(input *models.CreateProductRequest) (*models.Product, error) {
	product := models.Product{Name: input.Name, Price: input.Price, Quantity: input.Quantity}
	if err := ps.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (ps *productService) UpdateProduct(input *models.CreateProductRequest) (*models.Product, error) {
	var oldProduct models.Product
	if err := ps.db.First(&oldProduct, input.ID).Error; err != nil {
		return nil, errors.New("Record not found")
	}

	if err := ps.db.Model(&oldProduct).Updates(input).Error; err != nil {
		return nil, err
	}

	return &oldProduct, nil
}
