package envio

import (
	"gorm.io/gorm"
	"main.go/internal/core/domain"
)

type EnvioRepository struct {
	db *gorm.DB
}

func NewEnvioRepository(db *gorm.DB) *EnvioRepository {
	return &EnvioRepository{
		db: db,
	}
}

func (r *EnvioRepository) RegistrarEnvio(envio *domain.Envio) (uint, error) {
	result := r.db.Table("envios").Create(&envio)

	if result.Error != nil {
		return 0, result.Error
	}

	return envio.ID, nil

}
