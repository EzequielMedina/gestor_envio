package estados

import (
	"time"
)

type EnTransito struct{}

func (e *EnTransito) Nombre() string {
	return "En transito"
}

func (e *EnTransito) SiguienteEstado() EstadoEnvio {
	return &EnCamino{}
}

func (e *EnTransito) GuardarHistorial(envioID uint) HistorialEstado {
	return HistorialEstado{
		Estado:      e.Nombre(),
		EnvioID:     envioID,
		Comentario:  "El envio se encuentra en transito",
		FechaCambio: time.Now().UTC(),
	}
}
