package envio

import "main.go/internal/core/domain"

func (s *EnvioService) ValidarNumeroSeguimiento(numeroPedido string) error {
	err := validarNumeroSeguimiento(numeroPedido)
	if err != nil {
		return err
	}
	err = s.Repo.ValidarNumeroSeguimiento(numeroPedido)
	if err != nil {
		return err
	}
	return nil
}

func validarNumeroSeguimiento(numeroPedido string) error {
	if numeroPedido == "" {
		return domain.ErrNumeroSeguimientoRequerido
	}
	return nil
}
