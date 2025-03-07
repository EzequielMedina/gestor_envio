package domain

import "main.go/internal/core/domain/estados"

type Envio struct {
	ID                uint                `gorm:"primaryKey;column:id" json:"id"`
	NumeroSeguimiento string              `gorm:"type:varchar(50);unique;not null;column:numero_seguimiento" json:"numero_seguimiento"`
	Remitente         string              `gorm:"type:varchar(100);not null;column:remitente" json:"remitente"`
	Destinatario      string              `gorm:"type:varchar(100);not null;column:destinatario" json:"destinatario"`
	DireccionDestino  string              `gorm:"type:text;not null;column:direccion_destino" json:"direccion_destino"`
	Peso              float64             `gorm:"type:decimal(10,2);not null;column:peso" json:"peso"`
	Estado            string              `gorm:"type:enum('Pendiente','En tránsito','En camino','Intento fallido','Entregado');default:'Pendiente';column:estado" json:"estado"`
	EstadoActual      estados.EstadoEnvio `gorm:"-" json:"estado_actual"`
}

func (e *Envio) CambiarEstado() {
	e.EstadoActual = e.EstadoActual.SiguienteEstado()
}

func (e *Envio) GuardarCambioEstado() estados.HistorialEstado {
	return e.EstadoActual.GuardarHistorial(e.ID)
}

func (e *Envio) ObtenerEstadoActualApartirDelEstado(estado string) {
	switch estado {
	case "En tránsito":
		e.EstadoActual = &estados.EnTransito{}
	case "En camino":
		e.EstadoActual = &estados.EnCamino{}
	case "Entregado":
		e.EstadoActual = &estados.Entregado{}
	default:
		e.EstadoActual = &estados.Pendiente{}
	}

}
