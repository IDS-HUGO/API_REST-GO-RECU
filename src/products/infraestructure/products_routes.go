package infraestructure

import (
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	CreateProductController  *CreateProductController
	GetProductsController    *GetProductsController
	UpdateProductController  *UpdateProductController
	DeleteProductController  *DeleteProductController
	GetProductByIdController *GetProductByIdController
}

func NewProductRoutes(
	cpc *CreateProductController,
	gpc *GetProductsController,
	upc *UpdateProductController,
	dpc *DeleteProductController,
	gpbc *GetProductByIdController,
) *ProductRoutes {
	return &ProductRoutes{
		CreateProductController:  cpc,
		GetProductsController:    gpc,
		UpdateProductController:  upc,
		DeleteProductController:  dpc,
		GetProductByIdController: gpbc,
	}
}

func (pr *ProductRoutes) SetupRoutes(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.POST("", pr.CreateProductController.Execute) // Eliminamos la barra "/"
		products.GET("", pr.GetProductsController.Execute)    // Eliminamos la barra "/"
		products.GET("/:id", pr.GetProductByIdController.Execute)
		products.PUT("/:id", pr.UpdateProductController.Execute)
		products.DELETE("/:id", pr.DeleteProductController.Execute)
	}
}
