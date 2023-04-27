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
