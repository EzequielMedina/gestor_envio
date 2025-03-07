package ports

import (
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/core/domain"
)

type EnvioService interface {
	RegistrarEnvio(envioResquests *requests.EnvioRequest) (string, error)
	ValidarNumeroSeguimiento(numeroPedido string) error
	EnvioByNumeroSeguimiento(numeroPedido string) (*domain.Envio, error)
	ActualizarEnvio(numeroSeguimiento string, estado *requests.ActualizarEstadoRequest) error
}

type EnvioRepository interface {
	RegistrarEnvio(envio *domain.Envio) (uint, error)
	ValidarNumeroSeguimiento(numeroPedido string) error
	EnvioByNumeroSeguimiento(numeroPedido string) (*domain.Envio, error)
	ActualizarEnvio(envio *domain.Envio) error
}
