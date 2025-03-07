package transportista

import "main.go/internal/core/domain"

func (t *TransportistaService) ObtenerTransportistaByEmail(email string) (domain.Transportista, error) {
	err := validarTransportistaByEmail(email)
	if err != nil {
		return domain.Transportista{}, err
	}
	tranportista, err := t.Repo.ObtenerTransportistaByEmail(email)
	if err != nil {
		return domain.Transportista{}, err
	}
	return tranportista, nil
}
