package envio

import (
	"github.com/google/uuid"
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
	"main.go/internal/core/ports"
)

type EnvioService struct {
	HistorialEstado ports.HistorialEstadoService
	Repo            ports.EnvioRepository
}

func NewEnvioService(repo ports.EnvioRepository, historialEstado ports.HistorialEstadoService) *EnvioService {
	return &EnvioService{
		HistorialEstado: historialEstado,
		Repo:            repo,
	}
}

func validarDatosEnvio(envioResquests *requests.EnvioRequest) error {
	if envioResquests.Destinatario == "" {
		return domain.ErrDestinatarioRequerido
	}
	if envioResquests.DireccionDestino == "" {
		return domain.ErrDireccionRequerida
	}
	if envioResquests.Remitente == "" {
		return domain.ErrRemitenteRequerido
	}
	if envioResquests.Peso == 0 {
		return domain.ErrPesoRequerido
	}
	return nil
}

func crearNuevoEnvio(envioResquests *requests.EnvioRequest) *domain.Envio {
	envio := domain.Envio{
		Destinatario:      envioResquests.Destinatario,
		DireccionDestino:  envioResquests.DireccionDestino,
		Remitente:         envioResquests.Remitente,
		Peso:              envioResquests.Peso,
		EstadoActual:      &estados.Pendiente{},
		NumeroSeguimiento: uuid.New().String(),
	}
	envio.Estado = envio.EstadoActual.Nombre()

	return &envio

}
