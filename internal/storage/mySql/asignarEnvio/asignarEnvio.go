package asignarenvio

import (
	"gorm.io/gorm"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
)

type AsignarEnvioRepository struct {
	DB *gorm.DB
}

func NewAsignarEnvioRepository(db *gorm.DB) *AsignarEnvioRepository {
	return &AsignarEnvioRepository{
		DB: db,
	}
}

func (r *AsignarEnvioRepository) AsignarEnvioTransportista(asiginarEnvio *domain.AsignacionEnvio) error {

	result := r.DB.Table("asignacion_envios").Create(&asiginarEnvio)

	if result.Error != nil {
		return domain.ErrInternal
	}

	return nil
}

func (r *AsignarEnvioRepository) ValidarPedidoSinAsignar(envioId uint) error {

	var asignacionEnvio domain.AsignacionEnvio

	result := r.DB.Table("asignacion_envios").Where("envio_id = ?", envioId).First(&asignacionEnvio)

	if result.RowsAffected > 0 {
		return domain.ErrPedidoYaAsignado
	}

	return nil
}

func (r *AsignarEnvioRepository) ValidarTransportistaSinPedidoAsignado(transportistaId uint) error {

	var asignacionEnvio domain.AsignacionEnvio
	var pendiente = estados.Pendiente{}

	// Query mejorada
	result := r.DB.
		Table("envios").
		Select("ae.*").
		Joins("JOIN asignacion_envios ae ON ae.envio_id = envios.id").
		Joins("JOIN transportista t ON t.id = ae.transportista_id").
		Where("ae.transportista_id = ? AND envios.estado = ?", transportistaId, pendiente.Nombre()).
		First(&asignacionEnvio)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return domain.ErrInternal
	}

	//si traer un registro es portque el transportista tiene un pedido pendiente

	if result.RowsAffected > 0 {
		return domain.ErrTransportistaConPedidoPendiente
	}

	return nil
}
