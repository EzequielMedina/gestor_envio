package transportista

import "main.go/internal/core/ports"

type TransportistaHandler struct {
	TransportistaService ports.TransportistaService
}

func NewTransportistaHandler(transportistaService ports.TransportistaService) *TransportistaHandler {
	return &TransportistaHandler{
		TransportistaService: transportistaService,
	}
}
