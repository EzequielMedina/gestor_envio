package transportista

import (
	"main.go/internal/core/domain"
	"main.go/internal/core/ports"
)

type TransportistaService struct {
	Repo ports.TransportistaRepository
}

func NewTransportistaService(repo ports.TransportistaRepository) *TransportistaService {
	return &TransportistaService{
		Repo: repo,
	}
}

func validarTransportistaByEmail(email string) error {

	if email == "" {
		return domain.ErrEmailRequerido
	}
	return nil
}
