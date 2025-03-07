package asignarEnvio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/adapter/handler/api/response"
)

func (h *AsignarEnvioHandler) AsignarPedidoTransportista(c *gin.Context) {
	var request requests.AsignarEnvio
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AsignarEnvioService.AsignarEnvioTransportista(&request)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	response.HandleSuccess(c, resp)

}
