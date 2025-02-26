package envio

import "main.go/internal/core/ports"

type EnvioHandler struct {
	EnvioService ports.EnvioService
}

func NewEnvioHandler(envioService ports.EnvioService) *EnvioHandler {
	return &EnvioHandler{
		EnvioService: envioService,
	}
}
