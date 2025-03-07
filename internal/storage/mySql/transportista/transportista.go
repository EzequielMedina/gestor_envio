package transportista

import (
	"gorm.io/gorm"
	"main.go/internal/core/domain"
)

type TransportistaRepository struct {
	Db *gorm.DB
}

func NewTransportistaRepository(db *gorm.DB) *TransportistaRepository {
	return &TransportistaRepository{
		Db: db,
	}
}

func (r *TransportistaRepository) ObtenerTransportistaByEmail(email string) (domain.Transportista, error) {

	var transportista domain.Transportista

	result := r.Db.Table("transportista").Where("email = ?", email).First(&transportista)

	if result.Error == gorm.ErrRecordNotFound {
		return domain.Transportista{}, domain.ErrTransportistaNoEncontrado
	}

	if result.Error != nil {
		return domain.Transportista{}, domain.ErrInternal
	}

	return transportista, nil
}
