package handlers

import (
	"net/http"

	"pos_project/interfaces"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	iProduct interfaces.IProductImpl
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		iProduct: interfaces.IProductImpl{},
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.iProduct.GetProducts(c)

	//  := h.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// func (h *ProductHandler) CreateProduct(c *gin.Context) {
// 	var input models.CreateProductRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	product, err := h.service.CreateProduct(&input)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": product})
// }

// func (h *ProductHandler) GetProductById(c *gin.Context) {
// 	product, err := h.service.GetProductById(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": product})
// }

// func (h *ProductHandler) UpdateProduct(c *gin.Context) {
// 	var input models.CreateProductRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	product, err := h.service.UpdateProduct(&input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": product})
// }
