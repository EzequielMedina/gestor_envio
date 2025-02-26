package estados

import (
	"time"
)

type Pendiente struct{}

func (p *Pendiente) Nombre() string {
	return "Pendiente"
}

func (p *Pendiente) SiguienteEstado() EstadoEnvio {
	return &EnTransito{}
}

func (p *Pendiente) GuardarHistorial(envioID uint) HistorialEstado {
	return HistorialEstado{
		Estado:      p.Nombre(),
		EnvioID:     envioID,
		Comentario:  "El envio se encuentra pendiente",
		FechaCambio: time.Now().UTC(),
	}
}
