package asignarEnvio

import "main.go/internal/core/ports"

type AsignarEnvioHandler struct {
	AsignarEnvioService ports.AsignarEnvioService
}

func NewAsignarEnvioHandler(asignarEnvioService ports.AsignarEnvioService) *AsignarEnvioHandler {
	return &AsignarEnvioHandler{
		AsignarEnvioService: asignarEnvioService,
	}
}
