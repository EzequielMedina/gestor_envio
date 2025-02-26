package estados

import (
	"time"
)

type Entregado struct{}

func (e *Entregado) Nombre() string {
	return "Entregado"
}

func (e *Entregado) SiguienteEstado() EstadoEnvio {
	return nil
}

func (e *Entregado) GuardarHistorial(envioID uint) HistorialEstado {
	return HistorialEstado{
		Estado:      e.Nombre(),
		EnvioID:     envioID,
		Comentario:  "El envio se encuentra entregado",
		FechaCambio: time.Now().UTC(),
	}
}
