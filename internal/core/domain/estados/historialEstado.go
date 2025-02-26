package estados

import "time"

type HistorialEstado struct {
	ID          uint      `gorm:"primaryKey;column:id" json:"id"`
	EnvioID     uint      `gorm:"not null;column:envio_id" json:"envio_id"` // Agregado el campo EnvioID
	Estado      string    `gorm:"type:enum('Pendiente','En tránsito','En camino','Intento fallido','Entregado');not null;column:estado" json:"estado"`
	Comentario  string    `gorm:"type:text;column:comentario" json:"comentario"`
	FechaCambio time.Time `gorm:"column:fecha_cambio" json:"fecha_cambio"` // Usando time.Time para manejo adecuado de fecha
}
