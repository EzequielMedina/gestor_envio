package envio

import "main.go/internal/core/domain"

func (s *EnvioService) EnvioByNumeroSeguimiento(numeroPedido string) (*domain.Envio, error) {
	err := validarNumeroSeguimiento(numeroPedido)
	if err != nil {
		return nil, err
	}
	envio, err := s.Repo.EnvioByNumeroSeguimiento(numeroPedido)
	if err != nil {
		return nil, err
	}
	return envio, nil
}
