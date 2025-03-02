package asignarEnvio

import (
	"main.go/internal/core/ports"
)

type AsignarEnvioService struct {
	TransportistaService   ports.TransportistaService
	EnvioService           ports.EnvioService
	HistorialEstadoService ports.HistorialEstadoService
	Repo                   ports.AsignarEnvioRepository
}

func NewAsignarEnvioService(historialEstadoService ports.HistorialEstadoService, transportistaService ports.TransportistaService, envioService ports.EnvioService, repo ports.AsignarEnvioRepository) *AsignarEnvioService {
	return &AsignarEnvioService{
		HistorialEstadoService: historialEstadoService,
		TransportistaService:   transportistaService,
		EnvioService:           envioService,
		Repo:                   repo,
	}
}
