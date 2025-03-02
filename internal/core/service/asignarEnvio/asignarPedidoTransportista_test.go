package asignarEnvio_test

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"main.go/internal/adapter/handler/api/requests"
	"main.go/internal/adapter/handler/api/response"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
	"main.go/internal/core/service/asignarEnvio"
	"main.go/mocks"
)

type asingarEnvioTestInput struct {
	request *requests.AsignarEnvio
}

type asignarEnvioTestOutput struct {
	asignacion domain.AsignacionEnvio
	response   response.AsignarEnvio
	err        error
}

func TestEnvioService_EnvioByNUmeroSeguimiento(t *testing.T) {

	envioDomain := &domain.Envio{
		Destinatario:      gofakeit.Name(),
		DireccionDestino:  gofakeit.Address().Address,
		Remitente:         gofakeit.Name(),
		Peso:              float64(gofakeit.Number(1, 1000)),
		EstadoActual:      &estados.Pendiente{},
		NumeroSeguimiento: string(gofakeit.UUID()),
		Estado:            "Pendiente",
	}

	testTable := map[string]struct {
		mocks             func(asignacionRepor *mocks.MockAsignarEnvioRepository, envioServ *mocks.MockEnvioService, input *asingarEnvioTestInput, output *asignarEnvioTestOutput)
		input             *asingarEnvioTestInput
		output            *asignarEnvioTestOutput
		assertionFunction func(subT *testing.T, output *asignarEnvioTestOutput, err error)
	}{
		"empty ValidarNumeroSeguimiento": {
			mocks: func(asignacionRepor *mocks.MockAsignarEnvioRepository, envioServ *mocks.MockEnvioService, input *asingarEnvioTestInput, output *asignarEnvioTestOutput) {
				envioServ.EXPECT().ValidarNumeroSeguimiento(input.request.NumeroSeguimiento).Return(domain.ErrNumeroSeguimientoRequerido)
			},
			input: &asingarEnvioTestInput{
				request: &requests.AsignarEnvio{
					NumeroSeguimiento:  string(gofakeit.UUID()),
					EmailTransportista: gofakeit.Email(),
				},
			},
			output: &asignarEnvioTestOutput{
				asignacion: domain.AsignacionEnvio{},
				response:   response.AsignarEnvio{},
				err:        domain.ErrNumeroSeguimientoRequerido,
			},
			assertionFunction: func(subT *testing.T, output *asignarEnvioTestOutput, err error) {
				assert.Equal(subT, output.err, err)
			},
		},
		"EnvioByNumeroSeguimiento": {
			mocks: func(asignacionRepor *mocks.MockAsignarEnvioRepository, envioServ *mocks.MockEnvioService, input *asingarEnvioTestInput, output *asignarEnvioTestOutput) {
				envioServ.EXPECT().ValidarNumeroSeguimiento(input.request.NumeroSeguimiento).Return(nil)
				envioServ.EXPECT().EnvioByNumeroSeguimiento(input.request.NumeroSeguimiento).Return(nil, domain.ErrNumeroDeSeguimientoNoEncontrado)
			},
			input: &asingarEnvioTestInput{
				request: &requests.AsignarEnvio{
					NumeroSeguimiento:  string(gofakeit.UUID()),
					EmailTransportista: gofakeit.Email(),
				},
			},
			output: &asignarEnvioTestOutput{
				asignacion: domain.AsignacionEnvio{},
				response:   response.AsignarEnvio{},
				err:        domain.ErrNumeroDeSeguimientoNoEncontrado,
			},
			assertionFunction: func(subT *testing.T, output *asignarEnvioTestOutput, err error) {
				assert.Equal(subT, output.err, err)
			},
		},
		"error  ValidarPedidoSinAsignar": {
			mocks: func(asignacionRepor *mocks.MockAsignarEnvioRepository, envioServ *mocks.MockEnvioService, input *asingarEnvioTestInput, output *asignarEnvioTestOutput) {
				envioServ.EXPECT().ValidarNumeroSeguimiento(input.request.NumeroSeguimiento).Return(nil)
				envioServ.EXPECT().EnvioByNumeroSeguimiento(input.request.NumeroSeguimiento).Return(envioDomain, nil)
				asignacionRepor.EXPECT().ValidarPedidoSinAsignar(uint(envioDomain.ID)).Return(domain.ErrPedidoYaAsignado)
			},
			input: &asingarEnvioTestInput{
				request: &requests.AsignarEnvio{
					NumeroSeguimiento:  string(gofakeit.UUID()),
					EmailTransportista: gofakeit.Email(),
				},
			},
			output: &asignarEnvioTestOutput{
				asignacion: domain.AsignacionEnvio{},
				response:   response.AsignarEnvio{},
				err:        domain.ErrPedidoYaAsignado,
			},
			assertionFunction: func(subT *testing.T, output *asignarEnvioTestOutput, err error) {
				assert.Equal(subT, output.err, err)
			},
		},
	}
	for testName, test := range testTable {
		t.Run(testName, func(subT *testing.T) {
			ctrl := gomock.NewController(subT)
			defer ctrl.Finish()

			asignacionRepor := mocks.NewMockAsignarEnvioRepository(ctrl)
			envioServ := mocks.NewMockEnvioService(ctrl)
			histoServ := mocks.NewMockHistorialEstadoService(ctrl)
			transportistaService := mocks.NewMockTransportistaService(ctrl)

			test.mocks(asignacionRepor, envioServ, test.input, test.output)

			asignacionService := asignarEnvio.NewAsignarEnvioService(histoServ, transportistaService, envioServ, asignacionRepor)

			_, err := asignacionService.AsignarEnvioTransportista(test.input.request)

			test.assertionFunction(subT, test.output, err)
		})
	}

}
