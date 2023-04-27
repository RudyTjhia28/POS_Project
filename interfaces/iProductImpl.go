package interfaces

import (
	"pos_project/models"
	"pos_project/services"
)

type IProductImpl struct {
	productServices *services.ProductServices
}

func NewIProductImpl(service *services.ProductService) services.ProductService {
	return &services.ProductServices{}
}

func (ser *IProductImpl) GetProducts() (*[]models.Product, error) {
	res, err := ser.productServices.GetProducts()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ser *IProductImpl) GetProductById(id string) (*models.Product, error) {
	res, err := ser.productServices.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ser *IProductImpl) CreateProduct(req *models.CreateProductRequest) (*models.Product, error) {
	res, err := ser.productServices.CreateProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ser *IProductImpl) UpdateProduct(req *models.CreateProductRequest) (*models.Product, error) {
	res, err := ser.productServices.UpdateProduct(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
