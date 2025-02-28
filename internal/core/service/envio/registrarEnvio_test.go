package envio_test

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
	"main.go/internal/core/service/envio"
	"main.go/mocks"
)

type registrarEnvioTestedInput struct {
	EnvioRequest *requests.EnvioRequest
}

type registrarEnvioTestedOutput struct {
	numeroSeguimiento string
	err               error
	envioID           uint
}

func TestEnvioService_RegistrarEnvio(t *testing.T) {

	numeroSeguimiento := string(gofakeit.UUID())

	envioDomain := &domain.Envio{
		Destinatario:      gofakeit.Name(),
		DireccionDestino:  gofakeit.Address().Address,
		Remitente:         gofakeit.Name(),
		Peso:              float64(gofakeit.Number(1, 1000)),
		EstadoActual:      &estados.Pendiente{},
		NumeroSeguimiento: numeroSeguimiento,
		Estado:            "Pendiente",
	}

	testTable := map[string]struct {
		mocks             func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado)
		input             *registrarEnvioTestedInput
		output            *registrarEnvioTestedOutput
		inputHistorial    *estados.HistorialEstado
		assertionFunction func(subT *testing.T, output *registrarEnvioTestedOutput, err error)
	}{
		"empty remitente": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado) {

			},
			input: &registrarEnvioTestedInput{
				EnvioRequest: &requests.EnvioRequest{
					Remitente:        "",
					Destinatario:     gofakeit.Name(),
					DireccionDestino: gofakeit.Address().Address,
					Peso:             float64(gofakeit.Number(1, 1000)),
				},
			},
			output: &registrarEnvioTestedOutput{
				numeroSeguimiento: "",
				err:               domain.ErrRemitenteRequerido,
			},
			inputHistorial: nil,
			assertionFunction: func(subT *testing.T, output *registrarEnvioTestedOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Equal(subT, output.numeroSeguimiento, "")
				assert.NotNil(subT, err)
			},
		},
		"empty Destinatario": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado) {

			},
			input: &registrarEnvioTestedInput{
				EnvioRequest: &requests.EnvioRequest{
					Remitente:        gofakeit.Name(),
					Destinatario:     "",
					DireccionDestino: gofakeit.Address().Address,
					Peso:             float64(gofakeit.Number(1, 1000)),
				},
			},
			output: &registrarEnvioTestedOutput{
				numeroSeguimiento: "",
				err:               domain.ErrDestinatarioRequerido,
			}, inputHistorial: nil,

			assertionFunction: func(subT *testing.T, output *registrarEnvioTestedOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Equal(subT, output.numeroSeguimiento, "")
				assert.NotNil(subT, err)
			},
		},
		"empty Peso": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado) {

			},
			input: &registrarEnvioTestedInput{
				EnvioRequest: &requests.EnvioRequest{
					Remitente:        gofakeit.Name(),
					Destinatario:     gofakeit.Name(),
					DireccionDestino: gofakeit.Address().Address,
					Peso:             0,
				},
			},
			output: &registrarEnvioTestedOutput{
				numeroSeguimiento: "",
				err:               domain.ErrPesoRequerido,
			}, inputHistorial: nil,

			assertionFunction: func(subT *testing.T, output *registrarEnvioTestedOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Equal(subT, output.numeroSeguimiento, "")
				assert.NotNil(subT, err)
			},
		},
		"empty Direccion": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado) {

			},
			input: &registrarEnvioTestedInput{
				EnvioRequest: &requests.EnvioRequest{
					Remitente:        gofakeit.Name(),
					Destinatario:     gofakeit.Name(),
					DireccionDestino: "",
					Peso:             float64(gofakeit.Number(1, 1000)),
				},
			},
			output: &registrarEnvioTestedOutput{
				numeroSeguimiento: "",
				err:               domain.ErrDireccionRequerida,
			}, inputHistorial: nil,

			assertionFunction: func(subT *testing.T, output *registrarEnvioTestedOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Equal(subT, output.numeroSeguimiento, "")
				assert.NotNil(subT, err)
			},
		},
		"internal error repository": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado) {
				envioRepo.EXPECT().RegistrarEnvio(gomock.Any()).Return(uint(0), domain.ErrInternal)
			},
			input: &registrarEnvioTestedInput{
				EnvioRequest: &requests.EnvioRequest{
					Remitente:        gofakeit.Name(),
					Destinatario:     gofakeit.Name(),
					DireccionDestino: gofakeit.Address().Address,
					Peso:             float64(gofakeit.Number(1, 1000)),
				},
			},
			output: &registrarEnvioTestedOutput{
				numeroSeguimiento: "",
				err:               domain.ErrInternal,
			}, inputHistorial: nil,

			assertionFunction: func(subT *testing.T, output *registrarEnvioTestedOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Equal(subT, output.numeroSeguimiento, "")
				assert.NotNil(subT, err)
			},
		},
		"succes envio create repository": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado) {

				envioRepo.EXPECT().RegistrarEnvio(gomock.Any()).Return(envioDomain.ID, nil)
				historialEstado.EXPECT().RegistrarHistorialEstado(gomock.Any()).Return(nil)
			},
			input: &registrarEnvioTestedInput{
				EnvioRequest: &requests.EnvioRequest{
					Remitente:        envioDomain.Remitente,
					Destinatario:     envioDomain.Destinatario,
					DireccionDestino: envioDomain.DireccionDestino,
					Peso:             envioDomain.Peso,
				},
			},
			output: &registrarEnvioTestedOutput{
				numeroSeguimiento: envioDomain.NumeroSeguimiento,
				err:               nil,
				envioID:           envioDomain.ID,
			},
			inputHistorial: &estados.HistorialEstado{
				Estado:      envioDomain.Estado,
				EnvioID:     envioDomain.ID,
				Comentario:  "El envio se encuentra pendiente",
				FechaCambio: gofakeit.Date().UTC(),
			},
			assertionFunction: func(subT *testing.T, output *registrarEnvioTestedOutput, err error) {
				assert.Equal(subT, output.envioID, envioDomain.ID)
				assert.Nil(subT, err)
				assert.NotEmpty(subT, output.numeroSeguimiento)
			},
		},
		"succes envio create repository and fail historial estados empy envioId = 0": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, historialEstado *mocks.MockHistorialEstadoService, input *registrarEnvioTestedInput, output *registrarEnvioTestedOutput, inputHistorial *estados.HistorialEstado) {
				envioRepo.EXPECT().RegistrarEnvio(gomock.Any()).Return(envioDomain.ID, nil)
				historialEstado.EXPECT().RegistrarHistorialEstado(gomock.Any()).Return(domain.ErrEnvioIDRequerido)
			},
			input: &registrarEnvioTestedInput{
				EnvioRequest: &requests.EnvioRequest{
					Remitente:        envioDomain.Remitente,
					Destinatario:     envioDomain.Destinatario,
					DireccionDestino: envioDomain.DireccionDestino,
					Peso:             envioDomain.Peso,
				},
			},
			output: &registrarEnvioTestedOutput{
				numeroSeguimiento: envioDomain.NumeroSeguimiento,
				err:               nil,
				envioID:           envioDomain.ID,
			},
			inputHistorial: &estados.HistorialEstado{
				Estado:      envioDomain.Estado,
				EnvioID:     envioDomain.ID,
				Comentario:  "El envio se encuentra pendiente",
				FechaCambio: gofakeit.Date().UTC(),
			},
			assertionFunction: func(subT *testing.T, output *registrarEnvioTestedOutput, err error) {
				assert.Equal(subT, output.envioID, envioDomain.ID)
				assert.Equal(subT, output.err, err)
				assert.Empty(subT, output.numeroSeguimiento)
			},
		},
	}

	for testName, test := range testTable {

		t.Run(testName, func(subT *testing.T) {

			ctrl := gomock.NewController(subT)
			defer ctrl.Finish()

			envioRepo := mocks.NewMockEnvioRepository(ctrl)
			historialEstado := mocks.NewMockHistorialEstadoService(ctrl)

			test.mocks(envioRepo, historialEstado, test.input, test.output, test.inputHistorial)

			envioService := envio.NewEnvioService(envioRepo, historialEstado)

			numeroSeguimiento, err := envioService.RegistrarEnvio(test.input.EnvioRequest)
			output := &registrarEnvioTestedOutput{
				numeroSeguimiento: numeroSeguimiento,
				err:               err,
			}
			test.assertionFunction(subT, output, err)

		})
	}

}
