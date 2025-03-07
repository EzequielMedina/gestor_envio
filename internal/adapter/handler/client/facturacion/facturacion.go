package facturacion

import (
	"main.go/internal/adapter/config"
	"main.go/internal/adapter/handler/client"
)

type FacturarClient struct {
	Config            *config.ClientFacturacion
	clientFacturacion *client.ClientHttp
}

func NewFacturarClient(config *config.ClientFacturacion) *FacturarClient {
	return &FacturarClient{
		Config:            config,
		clientFacturacion: client.NewClientHttp(config.BaseUrl),
	}
}
