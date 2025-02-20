package infraestructure

import (
	"demo/src/Clients/applications"
	"demo/src/Clients/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateClientController struct {
	UseCase *applications.CreateClient
}

func NewCreateClientController(useCase *applications.CreateClient) *CreateClientController {
	return &CreateClientController{UseCase: useCase}
}

func (c *CreateClientController) Handle(ctx *gin.Context) {
	var client entities.Client
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos incompletos: " + err.Error(),
			"message": "Verifique los datos del cliente",
		})
		return
	}

	if err := c.UseCase.Execute(&client); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error de creación: " + err.Error(),
			"message": "No se pudo registrar el cliente",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Cliente creado exitosamente",
		"data":    client,
	})
}
