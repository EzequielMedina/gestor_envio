package estados

import "time"

type IntentoFallido struct{}

func (i IntentoFallido) Nombre() string {
	return "Intento fallido"
}

func (i IntentoFallido) SiguienteEstado() EstadoEnvio {
	return &IntentoFallido{}
}

func (i IntentoFallido) GuardarHistorial(envioID uint) HistorialEstado {
	return HistorialEstado{
		Estado:      i.Nombre(),
		EnvioID:     envioID,
		Comentario:  "El envio no pudo ser entregado",
		FechaCambio: time.Now().UTC(),
	}
}
