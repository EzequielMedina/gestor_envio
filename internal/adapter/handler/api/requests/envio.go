package requests

type EnvioRequest struct {
	Remitente        string  `json:"remitente"`
	Destinatario     string  `json:"destinatario"`
	DireccionDestino string  `json:"direccion_destino"`
	Peso             float64 `json:"peso"`
}
