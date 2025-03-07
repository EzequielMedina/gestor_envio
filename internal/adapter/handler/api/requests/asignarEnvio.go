package requests

type AsignarEnvio struct {
	NumeroSeguimiento  string `json:"numero_seguimiento"`
	EmailTransportista string `json:"email_transportista"`
}
