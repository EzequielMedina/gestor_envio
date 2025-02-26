package historialestado

import (
	"gorm.io/gorm"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
)

type HistorialEstadoRespository struct {
	db *gorm.DB
}

func NewHistorialEstadoRepository(db *gorm.DB) *HistorialEstadoRespository {
	return &HistorialEstadoRespository{
		db: db,
	}
}

func (r *HistorialEstadoRespository) RegistrarHistorialEstado(historialEstado *estados.HistorialEstado) error {

	result := r.db.Table("historial_estados").Create(&historialEstado)

	if result.Error != nil {
		return domain.ErrInternal
	}

	return nil

}
