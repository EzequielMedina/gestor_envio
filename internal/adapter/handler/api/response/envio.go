package response

type envioSeguimiento struct {
	NumeroSeguimiento string `json:"numero_seguimiento"`
}

func NewEnvioResponse(numeroSeguimiento string) *envioSeguimiento {
	return &envioSeguimiento{
		NumeroSeguimiento: numeroSeguimiento,
	}
}
