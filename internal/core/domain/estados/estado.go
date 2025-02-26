package estados

type EstadoEnvio interface {
	Nombre() string
	SiguienteEstado() EstadoEnvio
	GuardarHistorial(envioId uint) HistorialEstado
}
