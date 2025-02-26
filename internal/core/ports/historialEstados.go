package ports

import "main.go/internal/core/domain/estados"

type HistorialEstadoService interface {
	RegistrarHistorialEstado(historialEstado *estados.HistorialEstado) error
}

type HistorialEstadoRepository interface {
	RegistrarHistorialEstado(historialEstado *estados.HistorialEstado) error
}
