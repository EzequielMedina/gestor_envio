package envio

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/adapter/handler/api/response"
)

func (envioHandler *EnvioHandler) RegistrarEnvio(c *gin.Context) {
	var envioRequest requests.EnvioRequest
	if err := c.BindJSON(&envioRequest); err != nil {
		response.ValidationError(c, err)
		return
	}

	numeroSeguimiento, err := envioHandler.EnvioService.RegistrarEnvio(&envioRequest)

	if err != nil {
		response.HandleError(c, err)
		return
	}

	resp := response.NewEnvioResponse(numeroSeguimiento)
	response.HandleSuccess(c, resp)
}
