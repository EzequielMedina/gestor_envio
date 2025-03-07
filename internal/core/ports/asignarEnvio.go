package ports

import (
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/adapter/handler/api/response"
	"main.go/internal/core/domain"
)

type AsignarEnvioService interface {
	AsignarEnvioTransportista(asignarEnvioRequest *requests.AsignarEnvio) (response.AsignarEnvio, error)
}

type AsignarEnvioRepository interface {
	AsignarEnvioTransportista(asiginarEnvio *domain.AsignacionEnvio) error
	ValidarPedidoSinAsignar(envioId uint) error
	ValidarTransportistaSinPedidoAsignado(transportistaId uint) error
}
