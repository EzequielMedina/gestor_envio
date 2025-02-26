package historialEstado

import (
	"main.go/internal/core/domain/estados"
)

func (s *HistorialEstadoService) RegistrarHistorialEstado(historialEstado *estados.HistorialEstado) error {

	err := validarHistorialEstado(historialEstado)

	if err != nil {
		return err
	}

	err = s.Repo.RegistrarHistorialEstado(historialEstado)
	if err != nil {
		return err

	}

	return nil

}
