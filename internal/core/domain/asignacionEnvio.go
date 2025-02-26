package domain

import "time"

type AsignacionEnvio struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"id"`
	EnvioID         uint      `gorm:"not null;column:envio_id" json:"envio_id"` // Aquí agregamos EnvioID
	TransportistaID uint      `gorm:"not null;column:transportista_id" json:"transportista_id"`
	FechaAsignacion time.Time `gorm:"not null;column:fecha_asignacion" json:"fecha_asignacion"`
}
