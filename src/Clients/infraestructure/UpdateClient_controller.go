package infraestructure

import (
	"net/http"
	"strconv"

	"demo/src/Clients/applications"
	"demo/src/Clients/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateClientController struct {
	UseCase *applications.UpdateClient
}

func NewUpdateClientController(useCase *applications.UpdateClient) *UpdateClientController {
	return &UpdateClientController{UseCase: useCase}
}

func (c *UpdateClientController) Handle(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID inválido",
			"message": "El ID debe ser un número entero",
		})
		return
	}

	var updatedClient entities.Client
	if err := ctx.ShouldBindJSON(&updatedClient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos inválidos: " + err.Error(),
			"message": "Verifique los datos del cliente",
		})
		return
	}

	if err := c.UseCase.Execute(id, &updatedClient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al actualizar: " + err.Error(),
			"message": "No se pudo actualizar el cliente ID: " + idParam,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Cliente actualizado correctamente",
		"client_id": id,
		"data":      updatedClient,
	})
}
