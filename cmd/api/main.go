package main

import (
	"log"

	"main.go/internal/adapter/config"
	"main.go/internal/adapter/handler/api"
	EnvioHandler "main.go/internal/adapter/handler/api/envio"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
	envioService "main.go/internal/core/service/envio"
	hisotiralEstadoService "main.go/internal/core/service/historialEstado"
	mysql "main.go/internal/storage/mySql"
	"main.go/internal/storage/mySql/envio"
	historialestado "main.go/internal/storage/mySql/historialEstado"
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

	//dependency injection
	historialEstadoRepository := historialestado.NewHistorialEstadoRepository(db)
	historialEstadoService := hisotiralEstadoService.NewHistorialEstadoService(historialEstadoRepository)

	envioRepository := envio.NewEnvioRepository(db)
	envioServ := envioService.NewEnvioService(envioRepository, historialEstadoService)
	envioHand := EnvioHandler.NewEnvioHandler(envioServ)
	//router
	router, err := api.NewRouter(config.Http, *envioHand)
	if err != nil {
		log.Fatalf("Could not create the router: %v", err)
	}
	listenAddr := config.Http.Url + ":" + config.Http.Port

	err = router.Serve(listenAddr)
	if err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}
