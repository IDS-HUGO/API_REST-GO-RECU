package infraestructure

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}

type CreateProductController struct {
	cp application.CreateProduct
}

func NewCreateProductController(cp application.CreateProduct) *CreateProductController {
	return &CreateProductController{cp: cp}
}

func (cp_c *CreateProductController) Execute(c *gin.Context) {
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	product := entities.Product{
		Name:  req.Name,
		Price: req.Price,
	}

	err := cp_c.cp.Execute(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}
