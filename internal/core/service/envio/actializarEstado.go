package envio

import (
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
)

func (s *EnvioService) ActualizarEnvio(numeroSeguimiento string, estado *requests.ActualizarEstadoRequest) error {

	//validamos que el numero de seguimiento sea correcto

	err := s.ValidarNumeroSeguimiento(numeroSeguimiento)
	if err != nil {
		return err
	}

	//buscamos el envio por el numero de seguimiento

	envio, err := s.Repo.EnvioByNumeroSeguimiento(numeroSeguimiento)
	if err != nil {
		return err
	}

	nuevoEstado := estados.ObtenerEstadoActualApartirDelEstado(estado.Estado)

	envio.EstadoActual = nuevoEstado
	envio.Estado = nuevoEstado.Nombre()

	if envio.Estado != "En Transito" && envio.Estado != "Intento Fallido" {
		return domain.ErrEstadoInvalido
	}

	//actualizamos el estado del envio
	err = s.Repo.ActualizarEnvio(envio)
	if err != nil {
		return err
	}

	historial := envio.GuardarCambioEstado()

	if estado.Estado == "Intento Fallido" {
		historial.Comentario = estado.Razon
	}

	//guardamos el historial del estado

	err = s.HistorialEstado.RegistrarHistorialEstado(&historial)
	if err != nil {
		return err
	}

	return nil

}
