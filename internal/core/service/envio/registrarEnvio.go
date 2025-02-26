package envio

import (
	"main.go/internal/adapter/handler/api/requests"
)

func (s *EnvioService) RegistrarEnvio(envioResquests *requests.EnvioRequest) (string, error) {

	err := validarDatosEnvio(envioResquests)

	if err != nil {
		return "", err
	}

	envio := crearNuevoEnvio(envioResquests)

	_, err = s.Repo.RegistrarEnvio(envio)

	if err != nil {
		return "", err
	}
	historial := envio.GuardarCambioEstado()
	err = s.HistorialEstado.RegistrarHistorialEstado(&historial)

	if err != nil {
		return "", err
	}

	return envio.NumeroSeguimiento, nil

}
