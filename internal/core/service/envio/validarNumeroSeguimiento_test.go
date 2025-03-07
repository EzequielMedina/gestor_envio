package envio_test

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"main.go/internal/core/domain"
	"main.go/internal/core/service/envio"
	"main.go/mocks"
)

type validarByNumeroSeguimientoTestInput struct {
	numeroPedido string
}

type validarByNumeroSeguimientoTestOutput struct {
	err error
}

func TestEnvioService_ValidarByNUmeroSeguimiento(t *testing.T) {
	numeroSeguimiento := string(gofakeit.UUID())

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
				err: domain.ErrNumeroSeguimientoRequerido,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.err, err)
			},
		},
		"error in repository seguimiento no encontrado": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
				envioRepo.EXPECT().ValidarNumeroSeguimiento(input.numeroPedido).Return(domain.ErrNumeroDeSeguimientoNoEncontrado)
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: numeroSeguimiento,
			},
			output: &envioByNumeroSeguimientoTestOutput{
				err: domain.ErrNumeroDeSeguimientoNoEncontrado,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.err, err)
			},
		},
		"error in repository ": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
				envioRepo.EXPECT().ValidarNumeroSeguimiento(input.numeroPedido).Return(domain.ErrInternal)
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: numeroSeguimiento,
			},
			output: &envioByNumeroSeguimientoTestOutput{
				err: domain.ErrInternal,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.err, err)
			},
		},
		"error in repository pedido en otro estado ": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
				envioRepo.EXPECT().ValidarNumeroSeguimiento(input.numeroPedido).Return(domain.ErrPedidoNoEnProceso)
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: numeroSeguimiento,
			},
			output: &envioByNumeroSeguimientoTestOutput{
				err: domain.ErrPedidoNoEnProceso,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Equal(subT, output.err, err)
			},
		},
		"success": {
			mocks: func(envioRepo *mocks.MockEnvioRepository, input *envioByNumeroSeguimientoTestInput, output *envioByNumeroSeguimientoTestOutput) {
				envioRepo.EXPECT().ValidarNumeroSeguimiento(input.numeroPedido).Return(nil)
			},
			input: &envioByNumeroSeguimientoTestInput{
				numeroPedido: numeroSeguimiento,
			},
			output: &envioByNumeroSeguimientoTestOutput{
				err: nil,
			},
			assertionFunction: func(subT *testing.T, output *envioByNumeroSeguimientoTestOutput, err error) {
				assert.Nil(subT, output.err, err)
			},
		},
	}
	for testName, test := range testTable {
		t.Run(testName, func(subT *testing.T) {
			ctrl := gomock.NewController(subT)
			envioRepo := mocks.NewMockEnvioRepository(ctrl)

			test.mocks(envioRepo, test.input, test.output)

			s := envio.NewEnvioService(envioRepo, nil)

			err := s.ValidarNumeroSeguimiento(test.input.numeroPedido)

			test.assertionFunction(subT, test.output, err)
		})
	}

}
