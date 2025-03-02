package asignarEnvio

import (
	"time"

	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/adapter/handler/api/response"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
)

func (s *AsignarEnvioService) AsignarEnvioTransportista(asignarEnvioRequest *requests.AsignarEnvio) (response.AsignarEnvio, error) {

	//validamos que el numero del pedido sera valido
	err := s.EnvioService.ValidarNumeroSeguimiento(asignarEnvioRequest.NumeroSeguimiento)

	if err != nil {
		return response.AsignarEnvio{}, err
	}

	//validamos que el pedido no tenga un transportista asignado
	envio, err := s.EnvioService.EnvioByNumeroSeguimiento(asignarEnvioRequest.NumeroSeguimiento)
	if err != nil {
		return response.AsignarEnvio{}, err
	}
	err = s.Repo.ValidarPedidoSinAsignar(envio.ID)

	if err != nil {
		return response.AsignarEnvio{}, err
	}

	//validamos que el email pertenesca a un transportista
	transportista, err := s.TransportistaService.ObtenerTransportistaByEmail(asignarEnvioRequest.EmailTransportista)

	//validamos que el transportista no tenga un pedido asignado

	if err != nil {
		return response.AsignarEnvio{}, err
	}

	err = s.Repo.ValidarTransportistaSinPedidoAsignado(transportista.ID)

	if err != nil {
		return response.AsignarEnvio{}, err
	}

	//asignamos el pedido al transportista

	var asignacionEnvio = domain.AsignacionEnvio{
		EnvioID:         envio.ID,
		TransportistaID: transportista.ID,
		FechaAsignacion: time.Now().UTC(),
	}

	err = s.Repo.AsignarEnvioTransportista(&asignacionEnvio)

	//cambiar estado del pedido a en transito
	historial := actualizarEstadoPedido(envio)

	err = s.HistorialEstadoService.RegistrarHistorialEstado(historial)

	if err != nil {
		return response.AsignarEnvio{}, err
	}
	resp := response.NewAsignarEnvioResponse(asignacionEnvio.FechaAsignacion.String())
	//retornamos la respuesta

	return *resp, nil

}

func actualizarEstadoPedido(envio *domain.Envio) *estados.HistorialEstado {
	envio.ObtenerEstadoActualApartirDelEstado(envio.Estado)
	envio.EstadoActual.SiguienteEstado()
	envio.Estado = envio.EstadoActual.Nombre()
	historial := envio.GuardarCambioEstado()

	return &historial

}
