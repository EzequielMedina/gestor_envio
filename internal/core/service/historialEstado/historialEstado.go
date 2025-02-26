package historialEstado

import (
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
	"main.go/internal/core/ports"
)

type HistorialEstadoService struct {
	Repo ports.HistorialEstadoRepository
}

func NewHistorialEstadoService(repo ports.HistorialEstadoRepository) *HistorialEstadoService {
	return &HistorialEstadoService{
		Repo: repo,
	}
}

func validarHistorialEstado(historialEstado *estados.HistorialEstado) error {
	if historialEstado.Estado == "" {
		return domain.ErrEstadoRequerido
	}
	if historialEstado.EnvioID == 0 {
		return domain.ErrEnvioIDRequerido
	}
	return nil
}
