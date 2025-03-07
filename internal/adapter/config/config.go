package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App        *App
		DB         *DB
		Http       *Http
		ClientFact *ClientFacturacion
	}

	App struct {
		Name string
		Env  string
	}
	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
	Http struct {
		Env            string
		Url            string
		Port           string
		AllowedOrigins string
	}

	ClientFacturacion struct {
		BaseUrl        string
		GenerarFactura string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load("F:\\Eze\\go\\src\\gestor_envio\\.env")
		if err != nil {
			return nil, err
		}
	}
	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &Http{
		Env:            os.Getenv("APP_ENV"),
		Url:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("APP_ALLOWED_ORIGINS"),
	}
	clientFacturacion := &ClientFacturacion{
		BaseUrl:        os.Getenv("CLIENT_FACTURACION_BASE_URL"),
		GenerarFactura: os.Getenv("CLIENT_FACTURACION_GENERAR_FACTURA"),
	}
	return &Container{
		App:        app,
		DB:         db,
		Http:       http,
		ClientFact: clientFacturacion,
	}, nil
}
