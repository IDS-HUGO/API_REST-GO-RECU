package infraestructure

import (
	"net/http"
	"strconv"

	"demo/src/Clients/applications"

	"github.com/gin-gonic/gin"
)

type DeleteClientController struct {
	UseCase *applications.DeleteClient
}

func NewDeleteClientController(useCase *applications.DeleteClient) *DeleteClientController {
	return &DeleteClientController{UseCase: useCase}
}

func (c *DeleteClientController) Handle(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID inválido",
			"message": "El parámetro ID debe ser numérico",
		})
		return
	}

	if err := c.UseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al eliminar: " + err.Error(),
			"message": "No se pudo eliminar el cliente ID: " + idParam,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Cliente eliminado permanentemente",
		"deleted_id": id,
		"warning":    "Esta acción no se puede deshacer",
	})
}
