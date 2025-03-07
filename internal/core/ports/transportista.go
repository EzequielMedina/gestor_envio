package ports

import "main.go/internal/core/domain"

type TransportistaService interface {
	ObtenerTransportistaByEmail(email string) (domain.Transportista, error)
}

type TransportistaRepository interface {
	ObtenerTransportistaByEmail(email string) (domain.Transportista, error)
}
