package repositories

import (
	"pos_project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProducts(c *gin.Context) (*[]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

// func (r *ProductRepository) GetProductById(ctx *gin.Context, id string) (models.Product, error) {
// 	var product models.Product
// 	if err := r.db.First(&product, id).Error; err != nil {
// 		return models.Product{}, err
// 	}
// 	return product, nil
// }

// func (r *ProductRepository) CreateProduct(ctx *gin.Context, request models.CreateProductRequest) (models.Product, error) {
// 	product := models.Product{Name: request.Name, Price: request.Price, Quantity: request.Quantity}
// 	if err := r.db.Create(&product).Error; err != nil {
// 		return models.Product{}, err
// 	}
// 	return product, nil
// }

// func (r *ProductRepository) UpdateProduct(ctx *gin.Context, request models.CreateProductRequest) (models.Product, error) {
// 	var oldProduct models.Product
// 	if err := r.db.First(&oldProduct, request.ID).Error; err != nil {
// 		return models.Product{}, err
// 	}
// 	r.db.Model(&oldProduct).Updates(request)
// 	return oldProduct, nil
// }
