package estados

import (
	"time"
)

type EnCamino struct{}

func (e *EnCamino) Nombre() string {
	return "EnCamino"
}

func (e *EnCamino) SiguienteEstado() EstadoEnvio {
	return &Entregado{}
}

func (e *EnCamino) GuardarHistorial(envioID uint) HistorialEstado {
	return HistorialEstado{
		Estado:      e.Nombre(),
		EnvioID:     envioID,
		Comentario:  "El envio se encuentra en camino",
		FechaCambio: time.Now().UTC(),
	}
}
