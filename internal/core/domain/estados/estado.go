package estados

type EstadoEnvio interface {
	Nombre() string
	SiguienteEstado() EstadoEnvio
	GuardarHistorial(envioId uint) HistorialEstado
}

func ObtenerEstadoActualApartirDelEstado(estado string) EstadoEnvio {
	switch estado {
	case "Pendiente":
		return &Pendiente{}
	case "En transito":
		return &EnTransito{}
	case "En camino":
		return &EnCamino{}
	case "Entregado":
		return &Entregado{}
	case "Intento Fallido":
		return &IntentoFallido{}
	default:
		return nil
	}
}
