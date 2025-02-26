package ports

import (
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/core/domain"
)

type EnvioService interface {
	RegistrarEnvio(envioResquests *requests.EnvioRequest) (string, error)
}

type EnvioRepository interface {
	RegistrarEnvio(envio *domain.Envio) (uint, error)
}
