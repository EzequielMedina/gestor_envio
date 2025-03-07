package response

type AsignarEnvio struct {
	FechaAsignacion string `json:"fecha_asignacion"`
}

func NewAsignarEnvioResponse(fechaAsignacion string) *AsignarEnvio {
	return &AsignarEnvio{
		FechaAsignacion: fechaAsignacion,
	}
}
