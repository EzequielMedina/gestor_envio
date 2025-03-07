package main

import (
	"log"

	"main.go/internal/adapter/config"
	"main.go/internal/adapter/handler/api"
	asignarEnvioHandler "main.go/internal/adapter/handler/api/asignarEnvio"
	EnvioHandler "main.go/internal/adapter/handler/api/envio"
	transportistaHandler "main.go/internal/adapter/handler/api/transportista"
	"main.go/internal/adapter/handler/client/facturacion"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
	"main.go/internal/core/service/asignarEnvio"
	envioService "main.go/internal/core/service/envio"
	hisotiralEstadoService "main.go/internal/core/service/historialEstado"
	transportistaService "main.go/internal/core/service/transportista"
	mysql "main.go/internal/storage/mySql"
	asignarenvio "main.go/internal/storage/mySql/asignarEnvio"
	"main.go/internal/storage/mySql/envio"
	historialestado "main.go/internal/storage/mySql/historialEstado"
	"main.go/internal/storage/mySql/transportista"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(config.App.Name)

	db, err := mysql.Connect(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Transportista{}, &domain.Envio{}, &domain.AsignacionEnvio{}, &estados.HistorialEstado{})
	if err != nil {
		log.Fatal("Error al migrar las tablas:", err)
	}
	log.Println("Database connected")

	clienteFacturacion := facturacion.NewFacturarClient(config.ClientFact)

	//dependency injection
	historialEstadoRepository := historialestado.NewHistorialEstadoRepository(db)
	historialEstadoService := hisotiralEstadoService.NewHistorialEstadoService(historialEstadoRepository)

	envioRepository := envio.NewEnvioRepository(db)
	envioServ := envioService.NewEnvioService(envioRepository, historialEstadoService, *clienteFacturacion)
	envioHand := EnvioHandler.NewEnvioHandler(envioServ)

	transportistaRepo := transportista.NewTransportistaRepository(db)
	transportistaService := transportistaService.NewTransportistaService(transportistaRepo)
	transportistaHandler := transportistaHandler.NewTransportistaHandler(transportistaService)

	asignarEnvioRepository := asignarenvio.NewAsignarEnvioRepository(db)
	asingarEnvioService := asignarEnvio.NewAsignarEnvioService(historialEstadoService, transportistaService, envioServ, asignarEnvioRepository)
	asinarEnvioHandler := asignarEnvioHandler.NewAsignarEnvioHandler(asingarEnvioService)

	//router
	router, err := api.NewRouter(config.Http, *envioHand, *asinarEnvioHandler, *transportistaHandler)
	if err != nil {
		log.Fatalf("Could not create the router: %v", err)
	}
	listenAddr := config.Http.Url + ":" + config.Http.Port

	err = router.Serve(listenAddr)
	if err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}
