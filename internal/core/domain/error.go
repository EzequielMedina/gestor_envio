package domain

import (
	"errors"
)

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal error")
	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("record not found")
	// ErrNoUpdatedData is an error for when no data is provided to update
	ErrNoUpdatedData = errors.New("no data to update")
	// ErrConflictingData is an error for when data conflicts with existing data
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")
	// ErrInsufficientStock is an error for when product stock is not enough
	ErrInsufficientStock = errors.New("product stock is not enough")
	// ErrInsufficientPayment is an error for when total paid is less than total price
	ErrInsufficientPayment = errors.New("total paid is less than total price")
	// ErrTokenDuration is an error for when the token duration format is invalid
	ErrTokenDuration = errors.New("invalid token duration format")
	// ErrTokenCreation is an error for when the token creation fails
	ErrTokenCreation = errors.New("error creating token")
	// ErrExpiredToken is an error for when the access token is expired
	ErrExpiredToken = errors.New("access token has expired")
	// ErrInvalidToken is an error for when the access token is invalid
	ErrInvalidToken = errors.New("access token is invalid")
	// ErrInvalidCredentials is an error for when the credentials are invalid
	ErrInvalidCredentials = errors.New("invalid email or password")
	// ErrEmptyAuthorizationHeader is an error for when the authorization header is empty
	ErrEmptyAuthorizationHeader = errors.New("authorization header is not provided")
	// ErrInvalidAuthorizationHeader is an error for when the authorization header is invalid
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")
	// ErrInvalidAuthorizationType is an error for when the authorization type is invalid
	ErrInvalidAuthorizationType = errors.New("authorization type is not supported")
	// ErrUnauthorized is an error for when the user is unauthorized
	ErrUnauthorized = errors.New("user is unauthorized to access the resource")
	// ErrForbidden is an error for when the user is forbidden to access the resource
	ErrForbidden = errors.New("user is forbidden to access the resource")

	// ErrDestinatarioRequerido is an error for when the recipient is required
	ErrDestinatarioRequerido = errors.New("el destinatario es requerido")
	// ErrDireccionRequerida is an error for when the destination address is required
	ErrDireccionRequerida = errors.New("la dirección de destino es requerida")
	// ErrRemitenteRequerido is an error for when the sender is required
	ErrRemitenteRequerido = errors.New("el remitente es requerido")
	// ErrPesoRequerido is an error for when the weight is required
	ErrPesoRequerido = errors.New("el peso es requerido")

	// ErrInvalidEnvioID is an error for when the envio ID is invalid
	ErrNoSePudoRegistrarEnvio = errors.New("no se pudo registrar el envío")
	// ErrEnvioIDRequerido is an error for when the envio ID is required
	ErrEnvioIDRequerido = errors.New("el ID del envío es requerido")
	// ErrEstadoRequerido is an error for when the state is required
	ErrEstadoRequerido = errors.New("el estado es requerido")

	ErrNumeroSeguimientoRequerido = errors.New("el número de seguimiento es requerido")
	ErrPedidoNoEnProceso          = errors.New("el pedido no se encuentra en proceso")

	ErrNumeroDeSeguimientoNoEncontrado = errors.New("el número de seguimiento no fue encontrado")

	ErrPedidoNoExiste   = errors.New("el pedido no existe")
	ErrPedidoYaAsignado = errors.New("el pedido ya fue asignado")
	ErrEmailRequerido   = errors.New("el email es requerido")

	ErrTransportistaNoEncontrado       = errors.New("el transportista no fue encontrado")
	ErrEstadoInvalido                  = errors.New("el estado es inválido")
	ErrTransportistaConPedidoPendiente = errors.New("el transportista tiene un pedido pendiente")
)
