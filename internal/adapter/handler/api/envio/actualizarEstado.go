package envio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/adapter/handler/api/response"
)

func (e *EnvioHandler) ActualizarEstadoEnvio(c *gin.Context) {
	numeroSeguimiento := c.Param("numeroSeguimiento")
	var request requests.ActualizarEstadoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := e.EnvioService.ActualizarEnvio(numeroSeguimiento, &request)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.HandleSuccess(c, nil)
}
