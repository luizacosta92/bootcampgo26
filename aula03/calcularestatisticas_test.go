package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	t.Run("Minimum - retorna o menor valor", func(t *testing.T) {
		minFunc, err := operation("minimum")
		assert.NoError(t, err)
		assert.Equal(t, 2.0, minFunc(5, 8, 2, 10))
	})

	t.Run("Maximum - retorna o maior valor", func(t *testing.T) {
		maxFunc, err := operation("maximum")
		assert.NoError(t, err)
		assert.Equal(t, 10.0, maxFunc(5, 8, 2, 10))
	})

	t.Run("Average - retorna a média", func(t *testing.T) {
		avgFunc, err := operation("average")
		assert.NoError(t, err)
		assert.Equal(t, 6.25, avgFunc(5, 8, 2, 10))
	})

	t.Run("Operação inválida - retorna erro", func(t *testing.T) {
		_, err := operation("invalido")
		assert.Error(t, err)
		assert.Equal(t, "operação inválida", err.Error())
	})

	t.Run("Lista vazia - retorna 0", func(t *testing.T) {
		minFunc, _ := operation("minimum")
		maxFunc, _ := operation("maximum")
		avgFunc, _ := operation("average")

		assert.Equal(t, 0.0, minFunc())
		assert.Equal(t, 0.0, maxFunc())
		assert.Equal(t, 0.0, avgFunc())
	})
}
