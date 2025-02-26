package api

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/adapter/config"
	"main.go/internal/adapter/handler/api/envio"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.Http,
	envioHandler envio.EnvioHandler,

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
		}

	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
