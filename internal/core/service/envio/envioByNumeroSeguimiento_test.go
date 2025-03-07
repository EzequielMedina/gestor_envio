package envio_test

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"main.go/internal/core/domain"
	"main.go/internal/core/domain/estados"
	"main.go/internal/core/service/envio"
	"main.go/mocks"
)

type envioByNumeroSeguimientoTestInput struct {
	numeroPedido string
}

type envioByNumeroSeguimientoTestOutput struct {
	envio *domain.Envio
	err   error
}

func TestEnvioService_EnvioByNUmeroSeguimiento(t *testing.T) {
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
		mocks             func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput)
		input             *envioByNumeroSeguimientoTestInput
		output            *envioByNumeroSeguimientoTestOutput
		assertionFunction func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error)
	}{
		"empty numeroPedido": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: "",
			},
			output: &envioByNumeroSeguimientoTestOutput{
				envio: nil,
				err:   domain.ErrNumeroSeguimientoRequerido,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Nil(subT, output.envio)
			},
		},
		"error in repository seguimiento no encontrado": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
				envioRepo.EXPECT().EnvioByNumeroSeguimiento(input.numeroPedido).Return(nil, domain.ErrNumeroDeSeguimientoNoEncontrado)
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: numeroSeguimiento,
			},
			output: &envioByNumeroSeguimientoTestOutput{
				envio: nil,
				err:   domain.ErrNumeroDeSeguimientoNoEncontrado,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Nil(subT, output.envio)
			},
		},
		"error in repository ": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
				envioRepo.EXPECT().EnvioByNumeroSeguimiento(input.numeroPedido).Return(nil, domain.ErrInternal)
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: numeroSeguimiento,
			},
			output: &envioByNumeroSeguimientoTestOutput{
				envio: nil,
				err:   domain.ErrInternal,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.err, err)
				assert.Nil(subT, output.envio)
			},
		},
		"success": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
				envioRepo.EXPECT().EnvioByNumeroSeguimiento(input.numeroPedido).Return(envioDomain, nil)
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: numeroSeguimiento,
			},
			output: &envioByNumeroSeguimientoTestOutput{
				envio: envioDomain,
				err:   nil,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.envio, envioDomain)
				assert.Equal(subT, output.err, err)
			},
		},
	}
	for testName, test := range testTable {
		t.Run(testName, func(subT *testing.T) {
			ctrl := gomock.NewController(subT)
			envioRepo := mocks.NewMockEnvioRepository(ctrl)

			test.mocks(envioRepo, test.input, test.output)

			s := envio.NewEnvioService(envioRepo, nil)

			_, err := s.EnvioByNumeroSeguimiento(test.input.numeroPedido)

			test.assertionFunction(subT, test.output, err)
		})
	}

}
