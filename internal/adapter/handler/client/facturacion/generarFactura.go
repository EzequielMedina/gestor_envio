package facturacion

// metodo get donde le mandamos por la url el id del pedido

func (f *FacturarClient) GenerarFacturador(idPedido uint) (string, error) {
	//hacemos la peticion a otro microservicio
	url := f.Config.GenerarFactura + "?idPedido=" + string(rune(idPedido))
	resp, err := f.clientFacturacion.Get(url)

	if err != nil {
		return "", err
	}

	return resp, nil
}
