package infraestructure

import (
	"demo/src/Clients/applications"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetClientsController struct {
	UseCase *applications.GetClient
}

func NewGetClientsController(useCase *applications.GetClient) *GetClientsController {
	return &GetClientsController{UseCase: useCase}
}

func (c *GetClientsController) Handle(ctx *gin.Context) {
	clients, err := c.UseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener clientes: " + err.Error(),
			"message": "No se pudo recuperar la lista de clientes",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Clientes obtenidos exitosamente",
		"data":    clients,
		"count":   len(clients),
	})
}
