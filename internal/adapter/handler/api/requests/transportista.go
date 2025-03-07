package requests

type ActualizarEstadoRequest struct {
	Estado string `json:"estado" binding:"required"`
	Razon  string `json:"razon"`
}
