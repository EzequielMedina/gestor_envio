package envio

import (
	"gorm.io/gorm"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
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
		return uint(0), domain.ErrInternal
	}

	return envio.ID, nil

}

func (r *EnvioRepository) ValidarNumeroSeguimiento(numeroPedido string) error {
	//validamos que el numero de pedido solicitado este en proceso
	var envio domain.Envio

	result := r.db.Table("envios").Where("numero_seguimiento = ?", numeroPedido).First(&envio)

	if result.Error == gorm.ErrRecordNotFound {
		return domain.ErrNumeroDeSeguimientoNoEncontrado
	}

	if result.Error != nil {
		return domain.ErrInternal
	}
	var estadoPendiente estados.Pendiente
	if envio.Estado != estadoPendiente.Nombre() {
		return domain.ErrPedidoNoEnProceso
	}
	return nil
}

func (r *EnvioRepository) EnvioByNumeroSeguimiento(numeroPedido string) (*domain.Envio, error) {
	var envio domain.Envio
	result := r.db.Table("envios").Where("numero_seguimiento = ?", numeroPedido).First(&envio)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, domain.ErrNumeroDeSeguimientoNoEncontrado
	}
	if result.Error != nil {
		return nil, domain.ErrInternal
	}
	return &envio, nil
}

func (r *EnvioRepository) ActualizarEnvio(envio *domain.Envio) error {
	result := r.db.Table("envios").Where("id = ?", envio.ID).Updates(envio)
	if result.Error != nil {
		return domain.ErrInternal
	}
	return nil
}
