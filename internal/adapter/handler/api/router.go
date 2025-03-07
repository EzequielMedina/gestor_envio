package api

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/adapter/config"
	"main.go/internal/adapter/handler/api/asignarEnvio"
	"main.go/internal/adapter/handler/api/envio"
	"main.go/internal/adapter/handler/api/transportista"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.Http,
	envioHandler envio.EnvioHandler,
	asingarEnvioHandler asignarEnvio.AsignarEnvioHandler,
	transportistaHandler transportista.TransportistaHandler,

) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	v1 := router.Group("/v1")
	{
		envioGroup := v1.Group("/envio")
		{
			envioGroup.POST("/RegistrarEnvio", envioHandler.RegistrarEnvio)
			envioGroup.PATCH("/ActualizarEstado/:numeroSeguimiento/estado", envioHandler.ActualizarEstadoEnvio)
		}
		asignarEnvioGroup := v1.Group("/asignaciones")
		{
			asignarEnvioGroup.POST("/AsignarEnvio", asingarEnvioHandler.AsignarPedidoTransportista)
		}

	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
